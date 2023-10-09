package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"regexp"
	"server/internal/pkg/misc"
	"strings"
	"time"
)

type Config struct {
	GraphQL   string `yaml:"graphql"`
	Generated string `yaml:"generated"`
}

//go:generate go run generate.go

func main() {
	timestamp := time.Now().UnixNano()

	config, err := readConfig()
	if err != nil {
		fmt.Println("Error reading config:", err)
		return
	}

	fmt.Println("Starting with config:", config)

	files, err := os.ReadDir(config.GraphQL)
	if err != nil {
		fmt.Println("Error reading GraphQL directory:", err)
		return
	}

	fmt.Println("Found", len(files), "files")

	var rootSchema string
	var originalRootSchema string
	var specialTypes string

	for _, fileInfo := range files {
		if fileInfo.IsDir() || !strings.HasSuffix(fileInfo.Name(), ".graphql") || !strings.HasPrefix(fileInfo.Name(), "schema") {
			continue
		}

		graphqlContent, err := os.ReadFile(filepath.Join(config.GraphQL, fileInfo.Name()))
		if err != nil {
			fmt.Println("Error reading GraphQL file:", err)
			return
		}

		originalRootSchema = string(graphqlContent)
		rootSchema = generateRootSchema(string(graphqlContent))
	}

	for i := 0; i < len(files); i++ {
		fileInfo := files[i]
		if strings.HasPrefix(fileInfo.Name(), "schema") {
			files[i], files[len(files)-1] = files[len(files)-1], files[i]
			break
		}
	}

	for _, fileInfo := range files {
		if fileInfo.IsDir() || !strings.HasSuffix(fileInfo.Name(), ".graphql") {
			continue
		}

		goFileName := strings.TrimSuffix(fileInfo.Name(), ".graphql") + ".go"
		goFilePath := filepath.Join(config.Generated, goFileName)
		graphqlContent, err := os.ReadFile(filepath.Join(config.GraphQL, fileInfo.Name()))
		if err != nil {
			fmt.Println("Error reading GraphQL file:", err)
			return
		}
		if err != nil {
			fmt.Println("Error converting GraphQL file to schema:", err)
			return
		}

		content := "package generated\n\n// Code automatically generated; DO NOT EDIT.\n\n"

		if strings.HasPrefix(fileInfo.Name(), "schema") {
			content += rootSchema + fmt.Sprintf("\n\nfunc InitSchema() {\n%s}\n", specialTypes)
		} else {
			c, model, external := generateSchema(string(graphqlContent), originalRootSchema)
			specialTypes += external
			content += c

			modelFile, err := os.Create(filepath.Join(config.Generated, "models", strings.TrimSuffix(fileInfo.Name(), ".graphql")+"Models.go"))
			if err != nil {
				fmt.Println("Error creating model file:", err)
				return
			}

			_, err = modelFile.WriteString("package models\n\n// Code automatically generated; DO NOT EDIT.\n\n" + model)
			if err != nil {
				fmt.Println("Error writing to model file:", err)
				return
			}
		}

		goFile, err := os.Create(goFilePath)
		if err != nil {
			fmt.Println("Error creating Go file:", err)
			return
		}
		defer goFile.Close()

		_, err = goFile.WriteString(content)
		if err != nil {
			fmt.Println("Error writing to Go file:", err)
		}

	}
	fmt.Println("Finished operation in ", (time.Now().UnixNano()-timestamp)/1000, "ms")
}

func readConfig() (*Config, error) {
	configPath := "config.yml"

	configFile, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	decoder := yaml.NewDecoder(configFile)
	var config Config

	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

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

func generateRootSchema(schema string) string {
	lines := strings.Split(schema, "\n")
	roots := make([]string, 0)
	actionRegex := regexp.MustCompile(`^.+\(.+\): .+$`)

	var name string
	for i, line := range lines {
		if strings.HasPrefix(line, "type") {
			name = strings.Split(line, " ")[1]
			lines[i] = fmt.Sprintf("var root%s = graphql.NewObject(\n\tgraphql.ObjectConfig{\n\t\tName: \"Root%s\",\n\t\tFields: graphql.Fields{", name, name)
			roots = append(roots, fmt.Sprintf("\n\t\t%s: root%s,", name, name))
		} else if strings.HasPrefix(line, "}") {
			lines[i] = "\t\t},\n\t},\n)\n"
		} else if actionRegex.MatchString(line) {
			strings.Trim(line, " ")
			actionName := strings.Trim(strings.Split(line, "(")[0], " ")
			lines[i] = "\t\t\t\"" + actionName + "\": " + misc.FirstUpper(actionName) + name + ","
		}
	}

	return fmt.Sprintf("import \"github.com/graphql-go/graphql\"\n\n%svar RootSchema, _ = graphql.NewSchema(\n\tgraphql.SchemaConfig{%s\n\t},\n)", strings.Join(lines, "\n"), strings.Join(roots, ""))
}

func generateSchema(schema string, root string) (string, string, string) {
	objs := make([]string, 0)
	resolvers := make([]string, 0)
	models := make([]string, 0)
	newSchema := schemaToType(schema)
	newRoot := schemaToType(root)
	externalTypes := ""

	for _, t := range *newSchema {
		structs := []string{fmt.Sprintf("type %s struct {", t.Name)}
		types := []string{fmt.Sprintf("var %sType = graphql.NewObject(graphql.ObjectConfig{\n\tName: \"%s\", Fields: graphql.Fields{", misc.FirstLower(t.Name), t.Name)}
		mappers := []string{fmt.Sprintf("\nfunc ResultTo%s(result *neo4j.EagerResult) (*%s, error) {\n\tr, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], \"%s\")\n\tif err != nil {\n\t\treturn nil, err\n\t}\n", t.Name, t.Name, misc.FirstLower(t.Name)[:1])}
		returns := []string{fmt.Sprintf("\treturn &%s{", t.Name)}

		for _, field := range t.Fields {
			if !isGraphQLType(field.GraphQLType) {
				field.GoType = "*" + field.GoType
			}
			structs = append(structs, fmt.Sprintf("\t%s %s `json:\"%s\"`", field.GoName, field.GoType, field.JsonName))

			pointer := ""
			nonnullString := "graphql." + field.GraphQLType
			if !isGraphQLType(field.GraphQLType) {
				nonnullString = misc.FirstLower(field.GraphQLType) + "Type"
				pointer = "*"
			}
			if field.IsRequired {
				nonnullString = fmt.Sprintf("graphql.NewNonNull(%s)", nonnullString)
			}

			returns = append(returns, fmt.Sprintf("\t\t%s: %s%s,", field.GoName, pointer, strings.Replace(field.GraphQLName, "*", "&", -1)))
			if isGraphQLType(field.GraphQLType) {
				types = append(types, fmt.Sprintf("\t\t\t\"%s\": &graphql.Field{\n\t\t\t\tType: %s,\n\t\t\t},", field.GraphQLName, nonnullString))
				mappers = append(mappers, fmt.Sprintf("\t%s, err := neo4j.GetProperty[%s](r, \"%s\")\n\tif err != nil {\n\t\treturn nil, err\n\t}\n", field.GraphQLName, field.GoType, field.JsonName))
			} else {
				mappers = append(mappers, fmt.Sprintf("\t%s, err := ResultTo%s(result)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n", field.GraphQLName, field.GraphQLType))
				externalTypes += fmt.Sprintf("\t%sType.AddFieldConfig(\"%s\", &graphql.Field{Type: %s})\n", misc.FirstLower(t.Name), field.GraphQLName, nonnullString)
			}
		}

		for key := range *newRoot {
			for _, rootField := range (*newRoot)[key].Fields {
				if t.Name != rootField.GraphQLType {
					continue
				}

				res := []string{fmt.Sprintf("var %s = &graphql.Field{\n\tType: %sType,\n\tArgs: graphql.FieldConfigArgument{", rootField.GoName+key, misc.FirstLower(t.Name))}
				inputType := []string{fmt.Sprintf("type %sInput struct {", rootField.GoName)}
				inputMapper := []string{fmt.Sprintf("\t\tinput := &models.%sInput{", rootField.GoName)}

				for _, parameter := range rootField.Parameters {
					nonnullString := fmt.Sprintf("graphql.NewNonNull(graphql.%s)", parameter.Type)
					if !parameter.IsRequired {
						nonnullString = fmt.Sprintf("graphql.%s", parameter.Type)
					}
					res = append(res, fmt.Sprintf("\t\t\"%s\": &graphql.ArgumentConfig{\n\t\t\tType: %s,\n\t\t},", parameter.Name, nonnullString))
					inputType = append(inputType, fmt.Sprintf("\t%s %s `json:\"%s\"`", misc.FirstUpper(parameter.Name), parameter.GoType, parameter.Name))
					inputMapper = append(inputMapper, fmt.Sprintf("\t\t\t%s: p.Args[\"%s\"].(%s),", misc.FirstUpper(parameter.Name), parameter.Name, parameter.GoType))
				}

				res = append(res, fmt.Sprintf("\t},\n\tResolve: func(p graphql.ResolveParams) (interface{}, error) {\n%s\t\treturn resolvers.%s(p.Context, input)\n\t},", strings.Join(inputMapper, "\n")+"\t\t}\n\n", rootField.GoName+key))
				resolvers = append(resolvers, strings.Join(res, "\n")+"\n}\n")
				models = append(models, strings.Join(inputType, "\n")+"\n}\n")
			}
		}

		structs = append(structs, "}\n")
		types = append(types, "\t\t},\n\t},\n)\n")
		returns = append(returns, "\t}, nil\n}\n")

		models = append(models, strings.Join(structs, "\n")) // +strings.Join(mappers, "\n")+strings.Join(returns, "\n"))
		objs = append(objs, strings.Join(types, "\n"))
	}

	return "import (\n\t\"github.com/graphql-go/graphql\"\n\t\"server/internal/app/resolvers\"\n\t\"server/internal/pkg/generated/models\"\n)" +
		"\n\n" + strings.Join(objs, "\n") + strings.Join(resolvers, "\n"), strings.Join(models, "\n"), externalTypes
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
			goProperty := misc.FirstUpper(property)
			jsonProperty := misc.ConvertCamelToSnake(property)
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
				goProperty := misc.FirstUpper(property)
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
		return misc.FirstUpper(input)
	}
}

func isGraphQLType(input string) bool {
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
