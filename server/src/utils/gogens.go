package utils

import (
	"fmt"
	"regexp"
	"strings"
)

type Field struct {
	JsonName    string
	GoName      string
	GoType      string
	GraphQLType string
	GraphQLName string
	IsRequired  bool
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
	newSchema := schemaToType(schema)

	actions := make(map[string]string)

	for _, t := range newSchema {
		structs := make([]string, 0)
		types := make([]string, 0)
		maps := make([]string, 0)
		returns := make([]string, 0)

		structs = append(structs, fmt.Sprintf("type %s struct {", t.Name))
		types = append(types, fmt.Sprintf("var %sType = graphql.NewObject(\n\tgraphql.ObjectConfig{\n\t\tName: \"%s\",\n\t\tFields: graphql.Fields{", FirstLower(t.Name), t.Name))
		maps = append(maps, fmt.Sprintf("func ResultTo%s(result *neo4j.EagerResult) (*%s, error) {\n\taccountNode, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], \"a\")\n\tif err != nil {\n\t\treturn nil, err\n\t}\n", t.Name, t.Name))
		returns = append(returns, fmt.Sprintf("\treturn &%s{", t.Name))

		for _, field := range t.Fields {
			structs = append(structs, fmt.Sprintf("\t%s %s `json:\"%s\"`", field.GoName, field.GoType, field.JsonName))

			nonnullString := fmt.Sprintf("graphql.NewNonNull(graphql.%s)", field.GraphQLType)
			if !field.IsRequired {
				nonnullString = fmt.Sprintf("graphql.%s", field.GraphQLType)
			}
			types = append(types, fmt.Sprintf("\t\t\t\"%s\": &graphql.Field{\n\t\t\t\tType: %s,\n\t\t\t},", field.GraphQLName, nonnullString))
			maps = append(maps, fmt.Sprintf("\t%s, err := neo4j.GetProperty[%s](accountNode, \"%s\")\n\tif err != nil {\n\t\treturn nil, err\n\t}\n", field.GraphQLName, field.GoType, field.JsonName))
			returns = append(returns, fmt.Sprintf("\t\t%s: %s,", field.GoName, field.GraphQLName))
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

	//for _, returnType := range returnTypes {
	//	rootEqual := make(map[string]bool)
	//	isMutation := false
	//	for _, line := range strings.Split(root, "\n") {
	//		line = strings.Replace(line, "\r", "", -1)
	//		if strings.Contains(line, "type Mutation {") {
	//			isMutation = true
	//		} else if strings.Contains(line, "type Query {") {
	//			isMutation = false
	//		}
	//		if strings.HasSuffix(line, returnType) {
	//			rootEqual[line] = isMutation
	//		}
	//	}
	//
	//	for action, isMutation := range rootEqual {
	//		name := FirstUpper(strings.Trim(strings.Split(action, "(")[0], " "))
	//		args := make(map[string]string)
	//		var functionName string
	//		if isMutation {
	//			functionName = name + "Mutation"
	//		} else {
	//			functionName = name + "Query"
	//		}
	//
	//		re := regexp.MustCompile(`\w+\s*:\s*\w+!?`)
	//		matches := re.FindAllString(action, -1)
	//
	//		for _, match := range matches {
	//			parts := strings.SplitN(match, ":", 2)
	//			if len(parts) == 2 {
	//				argName := strings.TrimSpace(parts[0])
	//				argType := strings.TrimSpace(parts[1])
	//				args[argName] = argType
	//			}
	//		}
	//
	//		actions[action] = fmt.Sprintf("var %s = &graphql.Field{\n\tType: %sType,\n\tArgs: graphql.FieldConfigArgument{\n", functionName, FirstLower(returnType))
	//		for argName, argType := range args {
	//			isRequired := strings.Contains(argType, "!")
	//			if isRequired {
	//				argType = "graphql.NewNonNull(graphql." + strings.Replace(argType, "!", "", -1) + ")"
	//			} else {
	//				argType = "graphql." + argType
	//			}
	//			actions[action] = actions[action] + fmt.Sprintf("\t\t\"%s\": &graphql.ArgumentConfig{\n\t\t\tType: %s,\n\t\t},\n", argName, argType)
	//		}
	//
	//		actions[action] = actions[action] + "\t},\n\tResolve: func(p graphql.ResolveParams) (interface{}, error) {\n\t\treturn repository." + functionName + "(p)\n\t},\n}\n"
	//	}
	//}

	return "import (\n\t\"github.com/graphql-go/graphql\"\n\t\"github.com/neo4j/neo4j-go-driver/v5/neo4j\"\n\t\"server/src/db/repository\"\n)" +
		"\n\n" + strings.Join(objs, "\n") + "\n\n" + JoinMap(actions, "\n")
}

func schemaToType(schema string) map[string]Type {
	lines := strings.Split(schema, "\n")
	propertyRegex := regexp.MustCompile(`^.+.+: .+$`)
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
			graphqlType := strings.Replace(strings.Replace(strings.Trim(strings.Split(line, ":")[1], " "), "\r", "", -1), "!", "", -1)
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

	return res
}
