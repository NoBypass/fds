package middleware

import (
	"bytes"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Logger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","status":${status}, "error":"${error}",` +
			`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
			`"bytes_out":${bytes_out},${custom}}` + "\n",
		CustomTagFunc: func(c echo.Context, buf *bytes.Buffer) (int, error) {
			traceID := c.Get("traceID")
			msg := fmt.Sprintf("request: %s | %s (%d)", c.Request().RequestURI, c.Request().Method, c.Response().Status)
			send := fmt.Sprintf(`"trace_id":"%v","message":"%s"`, traceID, msg)
			return buf.WriteString(send)
		},
	})
}
