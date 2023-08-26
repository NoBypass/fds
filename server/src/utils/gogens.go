package utils

import (
	"fmt"
	"regexp"
	"strings"
)

type Parameter struct {
	Name       string
	Type       string
	IsRequired bool
}

type Field struct {
	JsonName    string
	GoName      string
	GoType      string
	GraphQLType string
	GraphQLName string
	IsRequired  bool
	Parameters  *[]Parameter
	IsMutation  *bool
}

type Type struct {
	Name   string
	Fields []Field
}

func GenerateRootSchema(schema string) string {
	lines := strings.Split(schema, "\n")
	actionRegex := regexp.MustCompile(`^.+\(.+\): .+$`)
	isMutation := false

	for i, line := range lines {
		if strings.HasPrefix(line, "type Query {") {
			lines[i] = "var rootQuery = graphql.NewObject(\n\tgraphql.ObjectConfig{\n\t\tName: \"RootQuery\",\n\t\tFields: graphql.Fields{"
			isMutation = false
		} else if strings.HasPrefix(line, "type Mutation {") {
			lines[i] = "var rootMutation = graphql.NewObject(\n\tgraphql.ObjectConfig{\n\t\tName: \"RootMutation\",\n\t\tFields: graphql.Fields{"
			isMutation = true
		} else if strings.HasPrefix(line, "}") {
			lines[i] = "\t\t},\n\t},\n)\n"
		} else if actionRegex.MatchString(line) {
			strings.Trim(line, " ")
			actionName := strings.Trim(strings.Split(line, "(")[0], " ")
			upper := FirstUpper(actionName)
			if isMutation {
				upper += "Mutation"
			} else {
				upper += "Query"
			}
			lines[i] = "\t\t\t\"" + actionName + "\": " + upper + ","
		}
	}

	return "import \"github.com/graphql-go/graphql\"\n\n" + strings.Join(lines, "\n") + "var RootSchema, _ = graphql.NewSchema(\n\tgraphql.SchemaConfig{\n\t\tQuery:    rootQuery,\n\t\tMutation: rootMutation,\n\t},\n)"
}

func GenerateSchema(schema string, root string) string {
	objs := make([]string, 0)
	resolvers := make([]string, 0)
	newSchema := schemaToType(schema)
	newRoot := schemaToType(root)

	for _, t := range newSchema {
		structs := []string{fmt.Sprintf("type %s struct {", t.Name)}
		types := []string{fmt.Sprintf("var %sType = graphql.NewObject(graphql.ObjectConfig{\n\tName: \"%s\", Fields: graphql.Fields{", t.Name, t.Name)}
		maps := []string{fmt.Sprintf("func ResultTo%s(r *neo4j.EagerResult) (*%s, error) {\n\tresult, _, err := neo4j.GetRecordValue[neo4j.Node](r.Records[0], \"%s\")\n\tif err != nil {\n\t\treturn nil, err\n\t}\n", t.Name, t.Name, FirstLower(t.Name)[0])}
		returns := []string{fmt.Sprintf("\treturn &%s{", t.Name)}

		for _, field := range t.Fields {
			structs = append(structs, fmt.Sprintf("\t%s %s `json:\"%s\"`", field.GoName, field.GoType, field.JsonName))

			nonnullString := fmt.Sprintf("graphql.NewNonNull(graphql.%s)", field.GraphQLType)
			if !field.IsRequired {
				nonnullString = fmt.Sprintf("graphql.%s", field.GraphQLType)
			}
			types = append(types, fmt.Sprintf("\t\t\t\"%s\": &graphql.Field{\n\t\t\t\tType: %s,\n\t\t\t},", field.GraphQLName, nonnullString))
			maps = append(maps, fmt.Sprintf("\t%s, err := neo4j.GetProperty[%s](result, \"%s\")\n\tif err != nil {\n\t\treturn nil, err\n\t}\n", field.GraphQLName, field.GoType, field.JsonName))
			returns = append(returns, fmt.Sprintf("\t\t%s: %s,", field.GoName, field.GraphQLName))
		}

		for key := range newRoot {
			for _, rootField := range newRoot[key].Fields {
				if t.Name != rootField.GraphQLType {
					continue
				}

				res := []string{fmt.Sprintf("var %s = &graphql.Field{\n\tType: %sType,\n\tArgs: graphql.FieldConfigArgument{", rootField.GoName+key, t.Name)}

				for _, parameter := range *rootField.Parameters {
					nonnullString := fmt.Sprintf("graphql.NewNonNull(graphql.%s)", parameter.Type)
					if !parameter.IsRequired {
						nonnullString = fmt.Sprintf("graphql.%s", parameter.Type)
					}
					res = append(res, fmt.Sprintf("\t\t\"%s\": &graphql.ArgumentConfig{\n\t\t\tType: %s,\n\t\t},", parameter.Name, nonnullString))
				}

				res = append(res, fmt.Sprintf("\t},\n\tResolve: func(p graphql.ResolveParams) (interface{}, error) {\n\t\treturn repository.%s(p), nil\n\t},", rootField.GoName+key))
				resolvers = append(resolvers, strings.Join(res, "\n")+"\n}\n")
			}
		}

		structs = append(structs, "}\n")
		types = append(types, "\t\t},\n\t},\n)\n")
		returns = append(returns, "\t}, nil\n}\n")

		var res string
		actions := [][]string{structs, types, maps, returns}
		for _, line := range actions {
			res += strings.Join(line, "\n") + "\n"
		}

		objs = append(objs, res)
	}

	return "import (\n\t\"github.com/graphql-go/graphql\"\n\t\"github.com/neo4j/neo4j-go-driver/v5/neo4j\"\n\t\"server/src/db/repository\"\n)" +
		"\n\n" + strings.Join(objs, "\n") + strings.Join(resolvers, "\n")
}

func schemaToType(schema string) map[string]Type {
	lines := strings.Split(schema, "\n")
	propertyRegex := regexp.MustCompile(`^.+.+: .+$`)
	resolverRegex := regexp.MustCompile(`^.+\(.+\): .+$`)
	res := make(map[string]Type)
	var isMutation bool
	var current string

	for _, line := range lines {
		if strings.HasPrefix(line, "type ") {
			current = strings.Split(line, " ")[1]
			res[current] = Type{
				Name: current,
			}

			if current == "Mutation" {
				isMutation = true
			} else if current == "Query" {
				isMutation = false
			}
		} else if propertyRegex.MatchString(line) {
			property := strings.Trim(strings.Split(line, ":")[0], " ")
			goProperty := FirstUpper(property)
			jsonProperty := convertCamelToSnake(property)
			isRequired := strings.Contains(line, "!")

			trim := strings.Split(line, " ")
			graphqlType := strings.Replace(strings.Replace(strings.Trim(trim[len(trim)-1], " "), "\r", "", -1), "!", "", -1)
			goType := strings.ToLower(graphqlType)

			if goType == "int" {
				goType = "int64"
			} else if goType == "float" {
				goType = "float64"
			}

			if strings.HasPrefix(property, "uuid") {
				property = "UUID"
			} else if strings.HasSuffix(property, "id") {
				property = property[:len(property)-2] + "ID"
			}

			if resolverRegex.MatchString(line) {
				paramString := strings.Split(strings.Split(line, "(")[1], ")")[0]
				property := strings.Trim(strings.Split(line, "(")[0], " ")
				goProperty := FirstUpper(property)
				params := strings.Split(paramString, ",")
				var parameters []Parameter
				for _, param := range params {
					param = strings.Trim(param, " ")
					paramName := strings.Split(param, ":")[0]
					paramType := strings.Split(param, ":")[1]
					parameters = append(parameters, Parameter{
						Name:       paramName,
						Type:       strings.Replace(strings.Trim(paramType, " "), "!", "", -1),
						IsRequired: strings.Contains(param, "!"),
					})
				}

				res[current] = Type{
					Name: res[current].Name,
					Fields: append(res[current].Fields, Field{
						JsonName:    jsonProperty,
						GoName:      goProperty,
						GoType:      goType,
						GraphQLType: graphqlType,
						GraphQLName: property,
						IsRequired:  isRequired,
						IsMutation:  &isMutation, // TODO incorrect
						Parameters:  &parameters,
					}),
				}
			} else {
				res[current] = Type{
					Name: res[current].Name,
					Fields: append(res[current].Fields, Field{
						JsonName:    jsonProperty,
						GoName:      goProperty,
						GoType:      goType,
						GraphQLType: graphqlType,
						GraphQLName: property,
						IsRequired:  isRequired,
					}),
				}
			}
		}
	}

	return res
}
