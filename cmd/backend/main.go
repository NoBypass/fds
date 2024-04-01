package main

import (
	"github.com/NoBypass/fds/internal/backend/auth"
	"github.com/NoBypass/fds/internal/backend/controller"
	"github.com/NoBypass/fds/internal/backend/middleware"
	"github.com/NoBypass/fds/internal/hypixel"
	"github.com/NoBypass/fds/internal/pkg/consts"
	"github.com/NoBypass/fds/internal/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

const VERSION = "v0.5.1"

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	closer := middleware.StartTracer(VERSION)
	defer closer.Close()

	println(`
   _______  ____  ____
  / __/ _ \/ __/ / __/__ _____  _____ ____
 / _// // /\ \  _\ \/ -_) __/ |/ / -_) __/
/_/ /____/___/ /___/\__/_/  |___/\__/_/   ` + consts.Purple.Sprint(VERSION) + `
Backend API for all FDS services written in ` + consts.WhiteOnCyan.Sprint(" GO ") + `
________________________________________________
`)

	e.HideBanner = true

	config := utils.ReadConfig()

	hypixelClient := hypixel.NewAPIClient(e)

	authService := auth.NewService(config.JWTSecret)
	discordController := controller.NewDiscordController(config, hypixelClient)

	e.Use(middleware.Recover())
	e.Use(middleware.Timeout())
	e.Use(middleware.Trace())
	e.Use(middleware.Logger())
	e.Use(middleware.Prepare(config))

	discord := e.Group("/discord")
	discord.Use(authService.DiscordAuthMiddleware())
	discord.POST("/verify", discordController.Verify)
	discord.GET("/member/:id", discordController.Member)
	discord.PATCH("/daily/:id", discordController.Daily)
	discord.POST("/bot-login", discordController.BotLogin)
	discord.DELETE("/revoke/:id", discordController.Revoke)
	discord.GET("/leaderboard/:page", discordController.Leaderboard)

	e.Logger.Fatal(e.Start(":8080"))
}
