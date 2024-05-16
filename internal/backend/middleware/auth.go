package middleware

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Auth(secret string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(secret),
		Skipper: func(c echo.Context) bool {
			path := c.Path()
			if path == "/discord/bot-login" {
				return true
			}

			return false
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
		},
	})
}
