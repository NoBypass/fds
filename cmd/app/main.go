package main

import (
	"github.com/NoBypass/fds/internal/app/auth"
	"github.com/NoBypass/fds/internal/app/controller"
	"github.com/NoBypass/fds/internal/app/middleware"
	"github.com/NoBypass/fds/internal/pkg/conf"
	"github.com/NoBypass/fds/internal/pkg/consts"
	"github.com/labstack/echo/v4"
)

const VERSION = "v0.3.0"

func main() {
	e := echo.New()

	println(`
   _______  ____  ____
  / __/ _ \/ __/ / __/__ _____  _____ ____
 / _// // /\ \  _\ \/ -_) __/ |/ / -_) __/
/_/ /____/___/ /___/\__/_/  |___/\__/_/   ` + consts.Purple.Sprint(VERSION) + `
Backend API for all FDS services written in ` + consts.WhiteOnCyan.Sprint(" GO ") + `
________________________________________________
`)

	e.HideBanner = true

	config := conf.ReadConfig()
	db := config.ConnectDB()

	authService := auth.NewService(config.JWTSecret)
	discordController := controller.NewDiscordController(db, config)

	e.Use(middleware.Logger())
	e.Use(middleware.Timeout())
	e.Use(middleware.Prepare(config))

	discord := e.Group("/discord")
	discord.Use(authService.DiscordAuthMiddleware())
	discord.POST("/verify", discordController.Verify)
	discord.PATCH("/daily/:id", discordController.Daily)
	discord.GET("/member/:id", discordController.Member)
	discord.GET("/leaderboard/:page", discordController.Leaderboard)
	discord.POST("/bot-login", discordController.BotLogin)

	e.Logger.Fatal(e.Start(":8080"))
}
