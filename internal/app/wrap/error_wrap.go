package wrap

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ErrorWrapper struct {
	c *echo.Context
}

func Error(c *echo.Context) ErrorWrapper {
	return ErrorWrapper{
		c,
	}
}

func (ew ErrorWrapper) NotFound(subj string) error {
	return (*ew.c).JSON(http.StatusNotFound, map[string]string{
		"message": fmt.Sprintf("%s not found", subj),
	})
}

func (ew ErrorWrapper) BadRequest(reason string) error {
	return (*ew.c).JSON(http.StatusBadRequest, map[string]string{
		"message": reason,
	})
}

func (ew ErrorWrapper) Success() error {
	return (*ew.c).JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}
