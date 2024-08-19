package env

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/gommon/log"
)

type Env struct {
	JaegerEndpoint string `env:"JAEGER_ENDPOINT"  env-required`

	HypixelAPIKey string `env:"HYPIXEL_API_KEY" env-required`

	BotPassword string `env:"BOT_PASSWORD"`
	JwtSecret   string `emv:"JWT_SECRET"`
	Port        string `env:"PORT" env-default:"1234"`

	SurrealPwd       string `env:"SURREAL_PWD" env-required`
	SurrealUser      string `env:"SURREAL_USER" env-required`
	SurrealHost      string `env:"SURREAL_HOST" env-required`
	SurrealNamespace string `env:"SURREAL_NAMESPACE" env-required`
	SurrealDB        string `env:"SURREAL_DB" env-required`

	Development string `env:"DEVELOPMENT"`
}

func Read() *Env {
	var cfg Env
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Fatalf("error reading evironnement variables: %s", err)
	}
	return &cfg
}
