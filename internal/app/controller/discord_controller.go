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

func NewDiscordController(db *surrealdb.DB) DiscordController {
	return &discordController{
		service: service.NewDiscordService(db),
	}
}

func (c *discordController) Verify(ctx echo.Context) error {
	errCh := make(chan error)
	defer close(errCh)

	inputCh := c.service.ParseVerify(ctx, errCh)
	memberCh := c.service.CreateMember(inputCh, errCh)

	select {
	case err := <-errCh:
		return err
	case member := <-memberCh:
		return ctx.JSON(http.StatusOK, member)
	}
}

func (c *discordController) Daily(ctx echo.Context) error {
	errCh := make(chan error)
	defer close(errCh)

	inputCh := c.service.ParseDaily(ctx, errCh)
	memberCh := c.service.GetMember(inputCh, errCh)
	xpCh := c.service.CheckDaily(memberCh, errCh)
	updatedMemberCh := c.service.GiveXP(memberCh, xpCh, errCh)

	select {
	case err := <-errCh:
		return err
	case member := <-updatedMemberCh:
		return ctx.JSON(http.StatusOK, member)
	}
}

func (c *discordController) BotLogin(ctx echo.Context) error {
	errCh := make(chan error)
	defer close(errCh)

	inputCh := c.service.ParseBotLogin(ctx, errCh)
	tokenCh := c.service.GetJWT(inputCh, ctx.Get("config").(*conf.Config), errCh)

	select {
	case err := <-errCh:
		return err
	case token := <-tokenCh:
		return ctx.JSON(http.StatusOK, map[string]string{
			"token": token,
		})
	}
}
