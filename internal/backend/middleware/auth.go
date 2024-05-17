package middleware

import (
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func Restrict(to model.AuthRole) echo.MiddlewareFunc {
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

			role, err := parseAudience(aud)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid role")
			}

			if role > to {
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
			}
			return next(c)
		}
	}
}

func Auth(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenStr := c.Request().Header.Get("Authorization")
			if tokenStr == "" {
				return next(c)
			}

			var claims jwt.RegisteredClaims
			token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (any, error) {
				return []byte(secret), nil
			})
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
			}

			c.Set("jwt", token)

			return next(c)
		}
	}
}

func parseAudience(aud []string) (model.AuthRole, error) {
	roles := make([]int, len(aud))
	for i, r := range aud {
		n, err := strconv.Atoi(r)
		if err != nil {
			return 0, err
		}

		roles[i] = n
	}

	smallest := roles[0]
	for _, r := range roles {
		if r < smallest {
			smallest = r
		}
	}
	return model.AuthRole(smallest), nil
}
