package controller

import (
	"github.com/NoBypass/fds/internal/backend/service"
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/labstack/echo/v4"
	"net/http"
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

func (c discordController) Member(ctx echo.Context) error {
	c.service.Request(ctx)

	id := ctx.Param("id")

	memberCh := c.service.GetMember(id)

	select {
	case err := <-c.service.Error():
		return err
	case member := <-memberCh:
		return ctx.JSON(http.StatusOK, member)
	}
}

func (c discordController) Verify(ctx echo.Context) error {
	c.service.Request(ctx)

	var input model.DiscordVerifyRequest
	err := ctx.Bind(&input)
	if err != nil {
		return err
	}

	playerResCh, memberBc := c.service.FetchHypixelPlayer(&input)
	playerBc, awaitVerify := c.service.VerifyHypixelSocials(memberBc.Attach(), playerResCh)
	c.service.PersistPlayer(playerBc.Attach())
	c.service.PersistMember(memberBc.Attach(), awaitVerify)
	c.service.RelateMemberToPlayer(memberBc.Attach(), playerBc.Attach())

	select {
	case err := <-c.service.Error():
		return err
	case actual := <-playerBc.Attach():
		return ctx.JSON(http.StatusOK, model.DiscordVerifyResponse{
			Actual: actual.Name,
		})
	}
}

func (c discordController) Revoke(ctx echo.Context) error {
	c.service.Request(ctx)

	id := ctx.Param("id")

	revokeCh := c.service.Revoke(id)

	select {
	case err := <-c.service.Error():
		return err
	case revokedMember := <-revokeCh:
		return ctx.JSON(http.StatusOK, revokedMember)
	}
}

func (c discordController) Daily(ctx echo.Context) error {
	c.service.Request(ctx)

	id := ctx.Param("id")

	memberCh := c.service.GetMember(id)
	updatedMemberCh := c.service.GiveDaily(memberCh)

	select {
	case err := <-c.service.Error():
		return err
	case member := <-updatedMemberCh:
		return ctx.JSON(http.StatusOK, member)
	}
}

func (c discordController) Leaderboard(ctx echo.Context) error {
	c.service.Request(ctx)

	page := ctx.Param("page")

	pageInt := c.service.StrToInt(page)
	leaderboardCh := c.service.GetLeaderboard(pageInt)

	select {
	case err := <-c.service.Error():
		return err
	case leaderboard := <-leaderboardCh:
		return ctx.JSON(http.StatusOK, leaderboard)
	}
}
