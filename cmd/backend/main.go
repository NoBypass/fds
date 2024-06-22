package main

import (
	"github.com/NoBypass/fds/internal/backend/controller"
	"github.com/NoBypass/fds/internal/backend/middleware"
	"github.com/NoBypass/fds/internal/backend/service"
	"github.com/NoBypass/fds/internal/external"
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/NoBypass/fds/internal/pkg/utils"
	"github.com/NoBypass/mincache"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/color"
	"github.com/labstack/gommon/log"
	"time"
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

	cache := mincache.New()
	e.Logger.Info("✓ Started cache")

	hypixelClient := external.NewHypixelAPIClient(cache, cfg.HypixelAPIKey)
	e.Logger.Info("✓ Connected to Hypixel API")

	e.Debug = cfg.Development != ""

	databaseSvc := service.NewDatabaseService(cfg)
	discordSvc := service.NewDiscordService(cfg, hypixelClient, databaseSvc)
	scrimsSvc := service.NewScrimsService(databaseSvc, cache)
	mojangSvc := service.NewMojangService(databaseSvc, cache)
	playerSvc := service.NewPlayerService(databaseSvc)
	authSvc := service.NewAuthService(cfg)
	minecraftSvc := service.NewMinecraftService(databaseSvc, scrimsSvc)

	scrimsController := controller.NewScrimsController(scrimsSvc, mojangSvc, minecraftSvc)
	playerController := controller.NewPlayerController(playerSvc, scrimsSvc)
	discordController := controller.NewDiscordController(discordSvc)
	authController := controller.NewAuthController(authSvc)

	mwc := middleware.NewCacheMiddleware(cache)

	e.Use(middleware.Timeout())
	e.Use(middleware.Logger())
	e.Use(middleware.Prepare(cfg))
	e.Use(middleware.Auth(cfg.JWTSecret))
	e.Use(middleware.AllowOrigin(cfg))
	e.Use(middleware.Trace())
	e.Use(middleware.Error())
	e.Use(middleware.Recover())

	discord := e.Group("/discord", middleware.Restrict(model.RoleBot))
	discord.POST("/verify", discordController.Verify)
	discord.GET("/member/:id", discordController.Member)
	discord.PATCH("/daily/:id", discordController.Daily)
	discord.DELETE("/revoke/:id", discordController.Revoke)
	discord.GET("/leaderboard/:page", discordController.Leaderboard)

	auth := e.Group("/auth")
	auth.POST("/bot", authController.Bot)

	player := e.Group("/player")
	player.GET("/exists/:name", playerController.Exists)

	scrims := player.Group("/scrims", mwc.Cache(5*time.Minute))
	scrims.GET("/:name", scrimsController.Player)
	scrims.GET("/:name/overview", scrimsController.Overview)
	//scrims.GET("/leaderboard/:page", scrimsController.Leaderboard)
	//scrims.GET("/scrim", )

	e.Logger.Fatal(e.Start(":8080"))
}
