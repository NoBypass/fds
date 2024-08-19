package controller

import (
	"github.com/NoBypass/fds/internal/app"
	"github.com/labstack/echo/v4"
	"net/http"
)

type DiscordController struct {
	uc *app.DiscordUseCase
}

func NewDiscordController(uc *app.DiscordUseCase) *DiscordController {
	return &DiscordController{
		uc: uc,
	}
}

func (ct *DiscordController) Auth(ctx echo.Context) error {
	var input inputPwd
	if err := ctx.Bind(&input); err != nil {
		return err
	}

	token, ok := ct.uc.Auth(ctx.Request().Context(), input.Pwd)
	if !ok {
		return echo.ErrUnauthorized
	}

	return ctx.JSON(http.StatusOK, token)
}
