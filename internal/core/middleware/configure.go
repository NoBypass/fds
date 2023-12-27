package middleware

import (
	"github.com/labstack/echo/v4"
	"server/internal/pkg/conf"
)

func Configure(ctx *conf.Context) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(ctx.Populate(c))
		}
	}
}
