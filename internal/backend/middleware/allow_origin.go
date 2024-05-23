package middleware

import (
	"github.com/NoBypass/fds/internal/pkg/utils"
	"github.com/labstack/echo/v4"
)

func AllowOrigin(cfg *utils.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if cfg.Development != "" {
				c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			}
			return next(c)
		}
	}
}
