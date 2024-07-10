package controller

import (
	"github.com/NoBypass/fds/internal/app"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ScrimsController struct {
	uc *app.MinecraftUseCase
}

func NewScrimsController(uc *app.MinecraftUseCase) *ScrimsController {
	return &ScrimsController{
		uc: uc,
	}
}

func (ct *ScrimsController) Leaderboard(ctx echo.Context) error {
	// TODO implement
	return nil
}

func (ct *ScrimsController) Player(ctx echo.Context) error {
	var input inputName
	if err := ctx.Bind(&input); err != nil {
		return err
	}

	player, err := ct.uc.GetScrimsPlayer(ctx.Request().Context(), input.Name)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, player)
}
