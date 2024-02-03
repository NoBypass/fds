package controller

import (
	"github.com/NoBypass/fds/internal/app/service"
	"github.com/NoBypass/fds/internal/pkg/conf"
	"github.com/labstack/echo/v4"
	"github.com/surrealdb/surrealdb.go"
	"net/http"
)

type DiscordController interface {
	Verify(c echo.Context) error
	Daily(c echo.Context) error
	BotLogin(c echo.Context) error
}

type discordController struct {
	service service.DiscordService
}

func NewDiscordController(db *surrealdb.DB, config *conf.Config) DiscordController {
	return &discordController{
		service.NewDiscordService(db, config),
	}
}

func (c discordController) Verify(ctx echo.Context) error {
	errCh := c.service.InjectErrorChan()

	inputCh := c.service.ParseVerify(ctx)
	mojangProfileCh, memberCh := c.service.FetchMojangProfile(inputCh)
	hypixelPlayerResCh, newMojangProfileCh := c.service.FetchHypixelPlayer(mojangProfileCh)
	verifiedMemberCh, hypixelPlayerCh := c.service.VerifyHypixelSocials(memberCh, hypixelPlayerResCh)
	done := c.service.Persist(newMojangProfileCh, verifiedMemberCh, hypixelPlayerCh)

	select {
	case err := <-errCh:
		return err
	case <-done:
		return ctx.JSON(http.StatusOK, map[string]bool{
			"success": true,
		})
	}
}

func (c discordController) Daily(ctx echo.Context) error {
	errCh := c.service.InjectErrorChan()

	inputCh := c.service.ParseDaily(ctx)
	memberCh := c.service.GetMember(inputCh) // TODO: use surql
	xpCh := c.service.CheckDaily(memberCh)
	updatedMemberCh := c.service.GiveXP(memberCh, xpCh)

	select {
	case err := <-errCh:
		return err
	case member := <-updatedMemberCh:
		return ctx.JSON(http.StatusOK, member)
	}
}

func (c discordController) BotLogin(ctx echo.Context) error {
	errCh := c.service.InjectErrorChan()

	inputCh := c.service.ParseBotLogin(ctx)
	tokenCh := c.service.GetJWT(inputCh)

	select {
	case err := <-errCh:
		return err
	case token := <-tokenCh:
		return ctx.JSON(http.StatusOK, map[string]string{
			"token": token,
		})
	}
}
