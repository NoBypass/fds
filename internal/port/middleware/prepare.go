package middleware

import (
	"github.com/NoBypass/fds/internal/common/env"
	"github.com/labstack/echo/v4"
)

func Prepare(config *env.Env) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("config", config)
			return next(c)
		}
	}
}
