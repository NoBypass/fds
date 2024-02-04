package conf

import (
	"os"
	"strconv"
)

type Config struct {
	HypixelAPIKey string
	DBNamespace   string
	JWTSecret     string
	BotPwd        string
	DBHost        string
	DBUser        string
	DBName        string
	DBPort        int
	DBPwd         string
	Port          int
}

func ReadConfig() *Config {
	port, err := strconv.Atoi(os.Getenv("port"))
	if err != nil {
		port = 8080
	}
	dbPort, err := strconv.Atoi(os.Getenv("db_port"))
	if err != nil {
		dbPort = 8000
	}
	return &Config{
		HypixelAPIKey: os.Getenv("hypixel_api_key"),
		DBNamespace:   os.Getenv("db_namespace"),
		JWTSecret:     os.Getenv("jwt_secret"),
		DBHost:        os.Getenv("db_host"),
		DBUser:        os.Getenv("db_user"),
		DBName:        os.Getenv("db_name"),
		BotPwd:        os.Getenv("bot_pwd"),
		DBPwd:         os.Getenv("db_pwd"),
		Port:          port,
		DBPort:        dbPort,
	}
}
