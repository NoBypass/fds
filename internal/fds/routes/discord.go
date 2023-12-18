package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Discord(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
