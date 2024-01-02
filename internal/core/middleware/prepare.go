package middleware

import (
	"github.com/labstack/echo/v4"
	"server/internal/pkg/conf"
)

func Prepare(config *conf.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("config", config)
			return next(c)
		}
	}
}
