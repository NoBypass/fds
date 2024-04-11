package main

import (
	"github.com/NoBypass/fds/internal/backend/auth"
	"github.com/NoBypass/fds/internal/backend/controller"
	"github.com/NoBypass/fds/internal/backend/database"
	"github.com/NoBypass/fds/internal/backend/middleware"
	"github.com/NoBypass/fds/internal/backend/service"
	"github.com/NoBypass/fds/internal/hypixel"
	"github.com/NoBypass/fds/internal/pkg/utils"
	"github.com/NoBypass/mincache"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/color"
	"github.com/labstack/gommon/log"
)

const VERSION = "v0.5.3"

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetLevel(log.INFO)

	e.Logger.Print(`:
   _______  ____  ____
  / __/ _ \/ __/ / __/__ _____  _____ ____
 / _// // /\ \  _\ \/ -_) __/ |/ / -_) __/
/_/ /____/___/ /___/\__/_/  |___/\__/_/   ` + color.Magenta(VERSION) + `
Backend API for all FDS services written in ` + color.CyanBg(color.White(" GO ")) + `
________________________________________________
`)

	closer := middleware.StartTracer(VERSION)
	defer closer.Close()
	e.Logger.Info("✓ Started tracer")

	cfg := utils.ReadConfig()
	e.Logger.Infof("✓ Loaded config %+v", cfg)

	db := database.Connect(cfg)
	e.Logger.Info("✓ Connected to SurrealDB")

	cache := mincache.New()
	e.Logger.Info("✓ Started cache")

	hypixelClient := hypixel.NewAPIClient(cache, cfg.HypixelAPIKey)
	e.Logger.Info("✓ Connected to Hypixel API")

	authService := auth.NewService(cfg.JWTSecret)
	discordSvc := service.NewDiscordService(cfg, hypixelClient, db)

	discordController := controller.NewDiscordController(discordSvc)

	e.Use(middleware.Recover())
	e.Use(middleware.Timeout())
	e.Use(middleware.Trace())
	e.Use(middleware.Logger())
	e.Use(middleware.Prepare(cfg))

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
