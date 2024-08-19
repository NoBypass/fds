package middleware

import (
	"github.com/NoBypass/fds/internal/common/env"
	"github.com/NoBypass/fds/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Restrict(to domain.AuthRole) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Get("jwt")
			if token == nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "no token provided")
			}

			claims := token.(*jwt.Token).Claims.(*jwt.RegisteredClaims)
			aud, err := claims.GetAudience()
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
			}

			can := domain.CanAccess(aud, to)
			if !can {
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
			}
			return next(c)
		}
	}
}

func Auth(secret *env.Env) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenStr := c.Request().Header.Get("Authorization")
			if tokenStr == "" {
				return next(c)
			}

			var claims jwt.RegisteredClaims
			token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (any, error) {
				return []byte(secret.JwtSecret), nil
			})
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
			}

			c.Set("jwt", token)

			return next(c)
		}
	}
}
