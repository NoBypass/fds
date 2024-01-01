package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"server/internal/pkg/consts"
)

func Logger() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:       true,
		LogURI:          true,
		LogMethod:       true,
		LogLatency:      true,
		LogResponseSize: true,
		BeforeNextFunc: func(c echo.Context) {
			c.Set("customValueFromContext", 42)
		},
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			str := fmt.Sprintf(" ⤷  %s  ", v.StartTime.Format("02/01/06 15:04:05"))
			switch {
			case v.Status >= 500:
				str += consts.Red.Sprint(v.Status)
			case v.Status >= 400:
				str += consts.Yellow.Sprint(v.Status)
			default:
				str += consts.Green.Sprint(v.Status)
			}

			str += " "

			switch v.Method {
			case "GET":
				str += consts.WhiteOnGreen.Sprint(" GET ")
			case "POST":
				str += consts.WhiteOnYellow.Sprint(" POST ")
			case "PUT":
				str += consts.WhiteOnBlue.Sprint(" PUT ")
			case "PATCH":
				str += consts.WhiteOnMagenta.Sprint(" PATCH ")
			case "DELETE":
				str += consts.WhiteOnRed.Sprint(" DELETE ")
			}
			str += fmt.Sprintf(" | Latency: %.2fms Size: %dB | %s ", float64(v.Latency)/1000000.0, v.ResponseSize, v.URI)
			fmt.Println(str)
			return nil
		},
	})
}
