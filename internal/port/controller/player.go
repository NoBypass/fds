package controller

import (
	"github.com/NoBypass/fds/internal/app"
	"github.com/labstack/echo/v4"
	"net/http"
)

type PlayerController struct {
	uc *app.MinecraftUseCase
}

func NewPlayerController(uc *app.MinecraftUseCase) *PlayerController {
	return &PlayerController{
		uc: uc,
	}
}

func (ct *PlayerController) Profile(ctx echo.Context) error {
	var input inputName
	if err := ctx.Bind(&input); err != nil {
		return err
	}

	player, err := ct.uc.GetMojangProfile(ctx.Request().Context(), input.Name)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, player)
}
