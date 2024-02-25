package controller

import (
	"github.com/NoBypass/fds/internal/app/service"
	"github.com/NoBypass/fds/internal/pkg/conf"
	"github.com/NoBypass/fds/internal/pkg/surreal_wrap"
	"github.com/NoBypass/fds/pkg/api"
	"github.com/labstack/echo/v4"
	"net/http"
)

type DiscordController interface {
	Verify(c echo.Context) error
	Daily(c echo.Context) error
	Member(c echo.Context) error
	Revoke(c echo.Context) error
	BotLogin(c echo.Context) error
	Leaderboard(c echo.Context) error
}

type discordController struct {
	service service.DiscordService
}

func NewDiscordController(db *surreal_wrap.DB, config *conf.Config) DiscordController {
	return &discordController{
		service.NewDiscordService(db, config),
	}
}

func (c discordController) Member(ctx echo.Context) error {
	errCh := c.service.InjectErrorChan()

	id := ctx.Param("id")

	memberCh := c.service.GetMember(id)

	select {
	case err := <-errCh:
		return err
	case member := <-memberCh:
		return ctx.JSON(http.StatusOK, member)
	}
}

func (c discordController) Verify(ctx echo.Context) error {
	errCh := c.service.InjectErrorChan()

	var input api.DiscordVerifyRequest
	err := ctx.Bind(&input)
	if err != nil {
		return err
	}

	verifiedCh := c.service.CheckIfAlreadyVerified(&input)
	mojangProfileCh, memberCh := c.service.FetchMojangProfile(verifiedCh)
	hypixelPlayerResCh, newMojangProfileCh := c.service.FetchHypixelPlayer(mojangProfileCh)
	verifiedMemberCh, hypixelPlayerCh := c.service.VerifyHypixelSocials(memberCh, hypixelPlayerResCh)
	actual := c.service.Persist(newMojangProfileCh, verifiedMemberCh, hypixelPlayerCh)

	select {
	case err := <-errCh:
		return err
	case actualName := <-actual:
		return ctx.JSON(http.StatusOK, api.DiscordVerifyResponse{
			Actual: actualName,
		})
	}
}

func (c discordController) Revoke(ctx echo.Context) error {
	errCh := c.service.InjectErrorChan()

	id := ctx.Param("id")

	revokeCh := c.service.Revoke(id)

	select {
	case err := <-errCh:
		return err
	case revokedMember := <-revokeCh:
		return ctx.JSON(http.StatusOK, revokedMember)
	}
}

func (c discordController) Daily(ctx echo.Context) error {
	errCh := c.service.InjectErrorChan()

	id := ctx.Param("id")

	memberCh := c.service.GetMember(id)
	xpCh := c.service.CheckDaily(memberCh)
	updatedMemberCh := c.service.GiveXP(memberCh, xpCh)

	select {
	case err := <-errCh:
		return err
	case member := <-updatedMemberCh:
		return ctx.JSON(http.StatusOK, member)
	}
}

func (c discordController) Leaderboard(ctx echo.Context) error {
	errCh := c.service.InjectErrorChan()

	page := ctx.Param("page")

	pageInt := c.service.StrToInt(page)
	leaderboardCh := c.service.GetLeaderboard(pageInt)

	select {
	case err := <-errCh:
		return err
	case leaderboard := <-leaderboardCh:
		return ctx.JSON(http.StatusOK, leaderboard)
	}
}

func (c discordController) BotLogin(ctx echo.Context) error {
	errCh := c.service.InjectErrorChan()

	var input api.DiscordBotLoginRequest
	err := ctx.Bind(&input)
	if err != nil {
		return err
	}

	tokenCh := c.service.GetJWT(&input)

	select {
	case err := <-errCh:
		return err
	case token := <-tokenCh:
		return ctx.JSON(http.StatusOK, api.DiscordBotLoginResponse{
			Token: token,
		})
	}
}
