package auth

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func (s *Service) DiscordAuthMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(s.secret),
		Skipper: func(c echo.Context) bool {
			path := c.Path()
			if path == "/discord/bot-login" {
				return true
			}

			return false
		},
	})
}
