package utils

import (
	"os"
	"strconv"
)

type Config struct {
	JaegerEndpoint string
	HypixelAPIKey  string
	DBNamespace    string
	Development    string
	JWTSecret      string
	BotPwd         string
	DBHost         string
	DBUser         string
	DBName         string
	DBPwd          string
	Port           int
}

func ReadConfig() *Config {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080
	}
	return &Config{
		JaegerEndpoint: os.Getenv("JAEGER_ENDPOINT"),
		HypixelAPIKey:  os.Getenv("HYPIXEL_API_KEY"),
		DBNamespace:    os.Getenv("DB_NAMESPACE"),
		Development:    os.Getenv("DEVELOPMENT"),
		JWTSecret:      os.Getenv("JWT_SECRET"),
		DBHost:         os.Getenv("DB_HOST"),
		DBUser:         os.Getenv("DB_USER"),
		DBName:         os.Getenv("DB_NAME"),
		BotPwd:         os.Getenv("BOT_PWD"),
		DBPwd:          os.Getenv("DB_PWD"),
		Port:           port,
	}
}
