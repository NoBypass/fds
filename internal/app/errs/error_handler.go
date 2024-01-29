package errs

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Handler(err error, c echo.Context) {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		_ = c.JSON(apiErr.Code, map[string]string{
			"error": err.Error(),
		})
	} else {
		_ = c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
}
