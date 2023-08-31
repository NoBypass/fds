package utils

import (
	"fmt"
	"regexp"
	"strings"
)

type Parameter struct {
	Name       string
	Type       string
	GoType     string
	IsRequired bool
}

type Field struct {
	JsonName    string
	GoName      string
	GoType      string
	GraphQLType string
	GraphQLName string
	IsRequired  bool
	Parameters  []Parameter
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

	for _, t := range *newSchema {
		structs := []string{fmt.Sprintf("type %s struct {", t.Name)}
		types := []string{fmt.Sprintf("var %sType = graphql.NewObject(graphql.ObjectConfig{\n\tName: \"%s\", Fields: graphql.Fields{", t.Name, t.Name)}
		maps := []string{fmt.Sprintf("func ResultTo%s(result *neo4j.EagerResult) (*%s, error) {\n\tr, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], \"%s\")\n\tif err != nil {\n\t\treturn nil, err\n\t}\n", t.Name, t.Name, FirstLower(t.Name)[:1])}
		returns := []string{fmt.Sprintf("\treturn &%s{", t.Name)}

		for _, field := range t.Fields {
			structs = append(structs, fmt.Sprintf("\t%s %s `json:\"%s\"`", field.GoName, field.GoType, field.JsonName))

			pointer := ""
			nonnullString := "graphql." + field.GraphQLType
			if !IsGraphQLType(field.GraphQLType) {
				nonnullString = field.GraphQLType + "Type"
				pointer = "*"
			}
			if field.IsRequired {
				nonnullString = fmt.Sprintf("graphql.NewNonNull(%s)", nonnullString)
			}
			types = append(types, fmt.Sprintf("\t\t\t\"%s\": &graphql.Field{\n\t\t\t\tType: %s,\n\t\t\t},", field.GraphQLName, nonnullString))
			returns = append(returns, fmt.Sprintf("\t\t%s: %s%s,", field.GoName, pointer, field.GraphQLName))
			if IsGraphQLType(field.GraphQLType) {
				maps = append(maps, fmt.Sprintf("\t%s, err := neo4j.GetProperty[%s](r, \"%s\")\n\tif err != nil {\n\t\treturn nil, err\n\t}\n", field.GraphQLName, field.GoType, field.JsonName))
			} else {
				maps = append(maps, fmt.Sprintf("\t%s, err := ResultTo%s(result)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n", field.GraphQLName, field.GraphQLType))
			}
		}

		for key := range *newRoot {
			for _, rootField := range (*newRoot)[key].Fields {
				if t.Name != rootField.GraphQLType {
					continue
				}

				res := []string{fmt.Sprintf("var %s = &graphql.Field{\n\tType: %sType,\n\tArgs: graphql.FieldConfigArgument{", rootField.GoName+key, t.Name)}
				inputType := []string{fmt.Sprintf("type %sInput struct {", rootField.GoName)}
				inputMapper := []string{fmt.Sprintf("\t\tinput := &%sInput{", rootField.GoName)}

				for _, parameter := range rootField.Parameters {
					nonnullString := fmt.Sprintf("graphql.NewNonNull(graphql.%s)", parameter.Type)
					if !parameter.IsRequired {
						nonnullString = fmt.Sprintf("graphql.%s", parameter.Type)
					}
					res = append(res, fmt.Sprintf("\t\t\"%s\": &graphql.ArgumentConfig{\n\t\t\tType: %s,\n\t\t},", parameter.Name, nonnullString))
					inputType = append(inputType, fmt.Sprintf("\t%s %s `json:\"%s\"`", FirstUpper(parameter.Name), parameter.GoType, parameter.Name))
					inputMapper = append(inputMapper, fmt.Sprintf("\t\t\t%s: p.Args[\"%s\"].(%s),", FirstUpper(parameter.Name), parameter.Name, parameter.GoType))
				}

				res = append(res, fmt.Sprintf("\t},\n\tResolve: func(p graphql.ResolveParams) (interface{}, error) {\n%s\t\treturn services.%s(p.Context, input), nil\n\t},", strings.Join(inputMapper, "\n")+"\t\t}\n\n", rootField.GoName+key))
				resolvers = append(resolvers, strings.Join(inputType, "\n")+"\n}\n\n"+strings.Join(res, "\n")+"\n}\n")
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

	return "import (\n\t\"github.com/graphql-go/graphql\"\n\t\"github.com/neo4j/neo4j-go-driver/v5/neo4j\"\n\t\"server/src/graph/services\"\n)" +
		"\n\n" + strings.Join(objs, "\n") + strings.Join(resolvers, "\n")
}

func schemaToType(schema string) *map[string]Type {
	lines := strings.Split(schema, "\n")
	propertyRegex := regexp.MustCompile(`^.+.+: .+$`)
	resolverRegex := regexp.MustCompile(`^.+\(.+\): .+$`)
	res := make(map[string]Type)
	var current string

	for _, line := range lines {
		if strings.HasPrefix(line, "type ") {
			current = strings.Split(line, " ")[1]
			res[current] = Type{
				Name: current,
			}

		} else if propertyRegex.MatchString(line) {
			property := strings.Trim(strings.Split(line, ":")[0], " ")
			goProperty := FirstUpper(property)
			jsonProperty := convertCamelToSnake(property)
			isRequired := strings.Contains(line, "!")

			trim := strings.Split(line, " ")
			graphqlType := strings.Replace(strings.Replace(strings.Trim(trim[len(trim)-1], " "), "\r", "", -1), "!", "", -1)
			goType := getGoType(graphqlType)

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
						GoType:     getGoType(paramType),
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
						Parameters:  parameters,
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

	return &res
}

func getGoType(input string) string {
	input = strings.Trim(strings.Replace(input, "!", "", -1), " ")
	switch strings.ToLower(input) {
	case "string":
		return "string"
	case "int":
		return "int64"
	case "float":
		return "float64"
	case "boolean":
		return "bool"
	case "id":
		return "string"
	default:
		return FirstUpper(input)
	}
}

func IsGraphQLType(input string) bool {
	switch strings.ToLower(input) {
	case "string":
		return true
	case "int":
		return true
	case "float":
		return true
	case "boolean":
		return true
	case "id":
		return true
	default:
		return false
	}
}