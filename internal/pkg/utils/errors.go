package utils

import (
	"github.com/NoBypass/surgo"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Error(res []surgo.Result) error {
	for _, r := range res {
		if r.Error != nil {
			return r.Error
		}
	}
	return nil
}

func ChannelNotOkError() error {
	return echo.NewHTTPError(http.StatusInternalServerError, "unexpectedly closed channel without having an error")
}
