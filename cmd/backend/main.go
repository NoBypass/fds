package main

import (
	"github.com/NoBypass/fds/internal/backend/controller"
	"github.com/NoBypass/fds/internal/backend/database"
	"github.com/NoBypass/fds/internal/backend/middleware"
	"github.com/NoBypass/fds/internal/backend/service"
	"github.com/NoBypass/fds/internal/external"
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/NoBypass/fds/internal/pkg/utils"
	"github.com/NoBypass/mincache"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/color"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetLevel(log.INFO)

	e.Logger.Print(`:
   _______  ____  ____
  / __/ _ \/ __/ / __/__ _____  _____ ____
 / _// // /\ \  _\ \/ -_) __/ |/ / -_) __/
/_/ /____/___/ /___/\__/_/  |___/\__/_/   ` + color.Magenta(utils.VERSION) + `
Backend API for all FDS services written in ` + color.CyanBg(color.White(" GO ")) + `
________________________________________________
`)

	closer := middleware.StartTracer()
	defer closer.Close()
	e.Logger.Info("✓ Started tracer")

	cfg := utils.ReadConfig()
	e.Logger.Info("✓ Loaded config")

	db := database.Connect(cfg)
	e.Logger.Info("✓ Connected to SurrealDB")

	cache := mincache.New()
	e.Logger.Info("✓ Started cache")

	hypixelClient := external.NewHypixelAPIClient(cache, cfg.HypixelAPIKey)
	e.Logger.Info("✓ Connected to Hypixel API")

	e.Debug = cfg.Development != ""

	discordSvc := service.NewDiscordService(cfg, hypixelClient, db)
	scrimsSvc := service.NewScrimsService(db, cache)
	authSvc := service.NewAuthService(cfg)

	discordController := controller.NewDiscordController(discordSvc)
	scrimsController := controller.NewScrimsController(scrimsSvc)
	authController := controller.NewAuthController(authSvc)

	e.Use(middleware.Timeout())
	e.Use(middleware.Trace())
	e.Use(middleware.Logger())
	e.Use(middleware.Prepare(cfg))
	e.Use(middleware.Auth(cfg.JWTSecret))

	discord := e.Group("/discord")
	discord.Use(middleware.Restrict(model.RoleBot))
	discord.POST("/verify", discordController.Verify)
	discord.GET("/member/:id", discordController.Member)
	discord.PATCH("/daily/:id", discordController.Daily)
	discord.DELETE("/revoke/:id", discordController.Revoke)
	discord.GET("/leaderboard/:page", discordController.Leaderboard)

	auth := e.Group("/auth")
	auth.POST("/bot", authController.Bot)

	scrims := e.Group("/scrims")
	//scrims.Use(middleware.Restrict(model.RoleMember))
	scrims.GET("/leaderboard/:page", scrimsController.Leaderboard)
	scrims.GET("/player/:name", scrimsController.Player)
	// scrims.GET("/scrim", )

	e.Logger.Fatal(e.Start(":8080"))
}
