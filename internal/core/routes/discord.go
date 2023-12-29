package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/surrealdb/surrealdb.go"
	"net/http"
	"server/internal/core/repository"
	"server/internal/core/services"
	"server/internal/pkg/conf"
	"server/internal/pkg/model"
)

func DiscordSignup(c echo.Context) error {
	ctx := conf.Preload(c)
	var input model.DiscordSignupInput
	err := c.Bind(&input)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid request query")
	}

	err = ctx.Repo.CreateDiscord(&input)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "success")
}

type dailyCtx struct {
	echo.Context
	Repo    repository.DiscordRepository
	Service services.DiscordService
}

func DiscordDaily(c echo.Context) error {
	ctx := conf.Preload(c)
	var input model.DiscordDailyInput
	err := c.Bind(&input)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid request query")
	}

	member, err := surrealdb.SmartUnmarshal[model.DiscordMember](ctx.Repo.Select("discord_member:" + input.ID))
	if err != nil {
		return err
	}

}
