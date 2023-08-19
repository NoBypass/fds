package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"server/src/utils"
	"strings"
)

type Config struct {
	GraphQL   string `yaml:"graphql"`
	Generated string `yaml:"generated"`
}

//go:generate go run generate.go

func main() {
	fmt.Println("Reading Config")

	config, err := readConfig()
	if err != nil {
		fmt.Println("Error reading config:", err)
		return
	}

	fmt.Println("Starting with config:", config)
	fmt.Println("Searching GraphQL directory for files")

	files, err := os.ReadDir(config.GraphQL)
	if err != nil {
		fmt.Println("Error reading GraphQL directory:", err)
		return
	}

	fmt.Println("Found", len(files), "files")
	fmt.Println("Finding root schema")

	var rootSchema string
	var originalRootSchema string
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
		rootSchema = utils.GenerateRootSchema(string(graphqlContent))
	}

	fmt.Println("Root schema found, generating files")

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
			content += rootSchema
		} else {
			content += utils.GenerateSchema(string(graphqlContent), originalRootSchema)
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

		fmt.Println("Generated file:", fileInfo.Name())
	}
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
