package utils

import (
	"github.com/NoBypass/surgo"
	"github.com/labstack/echo/v4"
)

func Error(res []surgo.Result) error {
	for _, r := range res {
		if r.Error != nil {
			return r.Error
		}
	}
	return nil
}

type (
	ErrMojangAPINotFound echo.HTTPError
)

func (e ErrMojangAPINotFound) Error() string {
	return e.Message.(string)
}
