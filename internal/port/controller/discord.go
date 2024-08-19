package controller

import (
	"github.com/NoBypass/fds/internal/app"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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

func (ct *DiscordController) Daily(ctx echo.Context) error {
	id := ctx.Param("id")

	daily, err := ct.uc.Daily(ctx.Request().Context(), id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, daily)
}

func (ct *DiscordController) Leaderboard(ctx echo.Context) error {
	p := ctx.QueryParam("page")
	page, err := strconv.Atoi(p)
	if err != nil {
		page = 0
	}

	leaderboard, err := ct.uc.GetLeaderboard(ctx.Request().Context(), page)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, leaderboard)
}

func (ct *DiscordController) Verify(ctx echo.Context) error {
	var input inputVerify
	if err := ctx.Bind(&input); err != nil {
		return err
	}

	ok, err := ct.uc.Verify(ctx.Request().Context(), input.DiscordID, input.IGN)
	if err != nil {
		return err
	} else if !ok {
		return echo.ErrUnauthorized
	}

	return ctx.NoContent(http.StatusOK)
}
