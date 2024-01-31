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
		Port:        port,
		DBHost:      os.Getenv("db_host"),
		DBPort:      dbPort,
		DBUser:      os.Getenv("db_user"),
		DBPwd:       os.Getenv("db_pwd"),
		DBNamespace: os.Getenv("db_namespace"),
		DBName:      os.Getenv("db_name"),
		JWTSecret:   os.Getenv("jwt_secret"),
		BotPwd:      os.Getenv("bot_pwd"),
	}
}
