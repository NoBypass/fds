package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"server/src/utils"
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
		rootSchema = utils.GenerateRootSchema(string(graphqlContent))
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
			c, model, external := utils.GenerateSchema(string(graphqlContent), originalRootSchema)
			specialTypes += external
			content += c

			// write file to generated/models containing the model
			modelFile, err := os.Create(filepath.Join(config.Generated, "models", strings.TrimSuffix(fileInfo.Name(), ".graphql")+"Models.go"))
			if err != nil {
				fmt.Println("Error creating model file:", err)
				return
			}

			_, err = modelFile.WriteString("package models\n\n// Code automatically generated; DO NOT EDIT.\n\n" + model)
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

		fmt.Println("Generated for file:", fileInfo.Name())
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
