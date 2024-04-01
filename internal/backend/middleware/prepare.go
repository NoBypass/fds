package middleware

import (
	"github.com/NoBypass/fds/internal/pkg/utils"
	"github.com/labstack/echo/v4"
)

func Prepare(config *utils.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("config", config)
			return next(c)
		}
	}
}
