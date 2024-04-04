package utils

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	JaegerEndpoint string
	HypixelAPIKey  string
	DBNamespace    string
	JWTSecret      string
	BotPwd         string
	DBHost         string
	DBUser         string
	DBName         string
	DBPort         int
	DBPwd          string
	Port           int
}

func ReadConfig() *Config {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080
	}
	cfg := &Config{
		JaegerEndpoint: os.Getenv("JAEGER_ENDPOINT"),
		HypixelAPIKey:  os.Getenv("HYPIXEL_API_KEY"),
		DBNamespace:    os.Getenv("DB_NAMESPACE"),
		JWTSecret:      os.Getenv("JWT_SECRET"),
		DBHost:         os.Getenv("DB_HOST"),
		DBUser:         os.Getenv("DB_USER"),
		DBName:         os.Getenv("DB_NAME"),
		BotPwd:         os.Getenv("BOT_PWD"),
		DBPwd:          os.Getenv("DB_PWD"),
		Port:           port,
	}

	fmt.Printf("Starting with config: %+v\n", cfg)
	return cfg
}
