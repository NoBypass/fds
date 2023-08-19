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
	// var actions string

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
				graphqlType = "graphql." + strings.Replace(graphqlType, "!", "", -1)
			}
			types[current] = types[current] + "\t\t\t\"" + jsonProperty + "\": &graphql.Field{\n\t\t\t\tType: " + graphqlType + ",\n\t\t\t},\n"
		}
	}

	return "import \"github.com/graphql-go/graphql\"\n" +
		"\n\n" + JoinMap(structs, "\n") +
		"\n\n" + JoinMap(types, "\n") +
		"\n\n"
}
