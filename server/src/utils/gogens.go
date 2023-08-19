package utils

import (
	"fmt"
	"regexp"
	"strings"
)

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
	lines := strings.Split(schema, "\n")
	propertyRegex := regexp.MustCompile(`^.+.+: .+$`)
	var current string
	var returnTypes []string
	structs := make(map[string]string)
	types := make(map[string]string)
	actions := make(map[string]string)

	for _, line := range lines {
		if strings.HasPrefix(line, "type ") {
			current = strings.Split(line, " ")[1]
			returnTypes = append(returnTypes, current)

			structs[current] = "type " + current + " struct {\n"
			types[current] = "var " + FirstLower(current) + "Type = graphql.NewObject(\n\tgraphql.ObjectConfig{\n\t\tName: \"" + current + "\",\n\t\tFields: graphql.Fields{\n"
		} else if strings.HasPrefix(line, "}") {
			structs[current] = structs[current] + "}\n"
			types[current] = types[current] + "\t\t},\n\t},\n)\n"
		} else if propertyRegex.MatchString(line) {
			property := strings.Trim(strings.Split(line, ":")[0], " ")
			goProperty := FirstUpper(property)
			jsonProperty := convertCamelToSnake(property)
			if strings.HasPrefix(property, "uuid") {
				property = "UUID"
			} else if strings.HasSuffix(property, "id") {
				property = property[:len(property)-2] + "ID"
			}

			graphqlType := strings.Replace(strings.Trim(strings.Split(line, ":")[1], " "), "\r", "", -1)
			goType := strings.Replace(strings.ToLower(graphqlType), "!", "", -1)
			isRequired := strings.Contains(graphqlType, "!")
			if isRequired {
				goType = "*" + goType
			}

			structs[current] = structs[current] + fmt.Sprintf("\t%s %s `json:\"%s\"`\n", goProperty, goType, jsonProperty)
			if isRequired {
				graphqlType = "graphql.NewNonNull(graphql." + strings.Replace(graphqlType, "!", "", -1) + ")"
			} else {
				graphqlType = "graphql." + graphqlType
			}
			types[current] = types[current] + "\t\t\t\"" + jsonProperty + "\": &graphql.Field{\n\t\t\t\tType: " + graphqlType + ",\n\t\t\t},\n"
		}
	}

	for _, returnType := range returnTypes {
		rootEqual := make(map[string]bool)
		isMutation := false
		for _, line := range strings.Split(root, "\n") {
			line = strings.Replace(line, "\r", "", -1)
			if strings.Contains(line, "type Mutation {") {
				isMutation = true
			} else if strings.Contains(line, "type Query {") {
				isMutation = false
			}
			if strings.HasSuffix(line, returnType) {
				rootEqual[line] = isMutation
			}
		}

		for action, isMutation := range rootEqual {
			name := FirstUpper(strings.Trim(strings.Split(action, "(")[0], " "))
			args := make(map[string]string)
			var functionName string
			if isMutation {
				functionName = name + "Mutation"
			} else {
				functionName = name + "Query"
			}

			re := regexp.MustCompile(`\w+\s*:\s*\w+!?`)
			matches := re.FindAllString(action, -1)

			for _, match := range matches {
				parts := strings.SplitN(match, ":", 2)
				if len(parts) == 2 {
					argName := strings.TrimSpace(parts[0])
					argType := strings.TrimSpace(parts[1])
					args[argName] = argType
				}
			}

			actions[action] = fmt.Sprintf("var %s = &graphql.Field{\n\tType: %sType,\n\tArgs: graphql.FieldConfigArgument{\n", functionName, FirstLower(returnType))
			for argName, argType := range args {
				isRequired := strings.Contains(argType, "!")
				if isRequired {
					argType = "graphql.NewNonNull(graphql." + strings.Replace(argType, "!", "", -1) + ")"
				} else {
					argType = "graphql." + argType
				}
				actions[action] = actions[action] + fmt.Sprintf("\t\t\"%s\": &graphql.ArgumentConfig{\n\t\t\tType: %s,\n\t\t},\n", argName, argType)
			}

			actions[action] = actions[action] + "\t},\n\tResolve: func(p graphql.ResolveParams) (interface{}, error) {\n\t\treturn repository." + functionName + "(p)\n\t},\n}\n"
		}
	}

	return "import (\n\t\"github.com/graphql-go/graphql\"\n\t\"server/src/db/repository\"\n)" +
		"\n\n" + JoinMap(structs, "\n") +
		"\n\n" + JoinMap(types, "\n") +
		"\n\n" + JoinMap(actions, "\n") +
		"\n\n"
}
