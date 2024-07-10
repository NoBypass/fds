package middleware

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func Error() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err != nil && !errors.Is(err, &echo.HTTPError{}) {
				ext.LogError(opentracing.SpanFromContext(c.Request().Context()), err)
			}
			return err
		}
	}
}
