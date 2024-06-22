package middleware

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"runtime"
)

func Recover() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					var err error
					switch x := r.(type) {
					case string:
						err = errors.New(x)
					case error:
						err = x
					default:
						err = fmt.Errorf("unknown panic: %v", r)
					}

					stack := make([]byte, 1<<10)
					length := runtime.Stack(stack, false)
					stack = stack[:length]
					msg := fmt.Sprintf("[PANIC RECOVER] %v %s\n", err, stack[:length])
					log.Errorf("recovered from panic: %v", msg)

					ext.LogError(opentracing.SpanFromContext(c.Request().Context()), fmt.Errorf(msg))

					c.Error(echo.NewHTTPError(500, "internal server error"))
				}
			}()
			return next(c)
		}
	}
}
