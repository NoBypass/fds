package conf

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func (c *Config) JWTConfig() *echojwt.Config {
	return &echojwt.Config{
		SigningKey: []byte(c.Authentication.Jwt.Secret),
		Skipper: func(c echo.Context) bool {
			path := c.Path()
			if path == "/discord/bot-login" {
				return true
			}

			return false
		},
	}
}
