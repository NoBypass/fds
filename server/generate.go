package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	GraphQL   string `yaml:"graphql"`
	Generated string `yaml:"generated"`
}

type Argument struct {
	Name     string
	Type     string
	Required bool
}

type Action struct {
	Name    string
	Args    []Argument
	Returns string
}

type Type struct {
	Name string
	Type string
}

type Schema struct {
	Mutations []Action
	Queries   []Action
	Types     map[string]Type
}

func stringToSchema(content string) (Schema, error) {
	// TODO
	return Schema{}, nil
}

//go:generate go run generate.go

func main() {
	config, err := readConfig()
	if err != nil {
		fmt.Println("Error reading config:", err)
		return
	}

	files, err := os.ReadDir(config.GraphQL)
	if err != nil {
		fmt.Println("Error reading GraphQL directory:", err)
		return
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
		graphqlContentSchema, err := stringToSchema(string(graphqlContent))
		if err != nil {
			fmt.Println("Error converting GraphQL file to schema:", err)
			return
		}

		content := "// Code automatically generated; DO NOT EDIT.\n\n"

		if strings.HasPrefix(fileInfo.Name(), "schema") {
			rootSchema, err := generateRootSchema(graphqlContentSchema)
			if err != nil {
				fmt.Println("Error generating root schema:", err)
				return
			}
			content += rootSchema
		} else {
			schema, err := generateSchema(graphqlContentSchema)
			if err != nil {
				fmt.Println("Error generating schema:", err)
				return
			}
			content += schema
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
}

func generateRootSchema(schema Schema) (string, error) {
	var res string = "package generated\n\nimport \"github.com/graphql-go/graphql\"\n\n"

	res += "var RootSchema, _ = graphql.NewSchema(\n\tgraphql.SchemaConfig{\n\t\tQuery:    rootQuery,\n\t\tMutation: rootMutation,\n\t},\n)"
	return res, nil
}

func generateSchema(schema Schema) (string, error) {
	// TODO
	return "", nil
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
