package controller

import (
	"github.com/NoBypass/fds/internal/backend/service"
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type DiscordController interface {
	Verify(c echo.Context) error
	Daily(c echo.Context) error
	Member(c echo.Context) error
	Revoke(c echo.Context) error
	Leaderboard(c echo.Context) error
}

type discordController struct {
	service service.DiscordService
}

func NewDiscordController(svc service.DiscordService) DiscordController {
	return &discordController{
		svc,
	}
}

func (ct discordController) Member(c echo.Context) error {
	id := c.Param("id")

	memberCh, err := ct.service.GetMember(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, memberCh)
}

func (ct discordController) Verify(c echo.Context) error {
	ctx := c.Request().Context()

	var input model.DiscordVerifyRequest
	err := c.Bind(&input)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	playerRes, member, err := ct.service.FetchHypixelPlayer(ctx, &input)
	if err != nil {
		return err
	}

	// TODO check socials with db and scrims as well
	player, err := ct.service.VerifyHypixelSocials(ctx, member, playerRes)
	if err != nil {
		return err
	}

	// TODO implement already exists check
	err = ct.service.PersistPlayer(ctx, player)
	if err != nil {
		return err
	}

	// TODO parallelize
	err = ct.service.PersistMember(ctx, member)
	if err != nil {
		return err
	}

	err = ct.service.RelateMemberToPlayer(ctx, member, player)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]any{
		"actual": player.Name,
	})
}

func (ct discordController) Revoke(c echo.Context) error {
	id := c.Param("id")

	member, err := ct.service.Revoke(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, member)
}

func (ct discordController) Daily(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	member, err := ct.service.GetMember(ctx, id)
	if err != nil {
		return err
	}

	err = ct.service.GiveDaily(ctx, member)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, member)
}

func (ct discordController) Leaderboard(c echo.Context) error {
	page := c.Param("page")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid page number")
	}

	member, err := ct.service.GetLeaderboard(c.Request().Context(), pageInt)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, member)
}
