package utils

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type ENV struct {
	HypixelKey string `yaml:"HypixelKey"`
	ClientURL  string `yaml:"ClientUrl"`
	JWTSecret  string `yaml:"JWTSecret"`
	DB         struct {
		URI      string `yaml:"URI"`
		Username string `yaml:"Username"`
		Password string `yaml:"Password"`
	} `yaml:"DB"`
}

func FetchEnv() ENV {
	file := "env.yml"
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	configuration := ENV{}
	if err := yaml.Unmarshal(data, &configuration); err != nil {
		log.Fatalf("Error unmarshaling YAML: %v", err)
	}

	return configuration
}
