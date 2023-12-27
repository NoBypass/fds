package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"server/internal/pkg/conf"
	"server/internal/pkg/model"
)

func DiscordSignup(c echo.Context) error {
	ctx := conf.Preload(c)
	var discordMember model.DiscordMemberInput
	err := c.Bind(&discordMember)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid request query")
	}

	err = ctx.Repo.CreateDiscord(&discordMember)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "success")
}
