package startup

import (
	"github.com/NoBypass/fds/internal/backend/controller"
	"github.com/NoBypass/fds/internal/backend/middleware"
	"github.com/NoBypass/fds/internal/backend/service"
	"github.com/NoBypass/fds/internal/frontend/renderer"
	"github.com/NoBypass/fds/internal/pkg/model"
)

func Router(app *App) {
	e := app.Echo

	static := renderer.New()
	e.Renderer = static

	discordSvc := service.NewDiscordService(app.Cfg, app.HypixelClient, app.DB)
	scrimsSvc := service.NewScrimsService(app.DB, app.Cache)
	mojangSvc := service.NewMojangService(app.DB, app.Cache)
	playerSvc := service.NewPlayerService(app.DB)
	authSvc := service.NewAuthService(app.Cfg)

	scrimsController := controller.NewScrimsController(scrimsSvc, mojangSvc)
	playerController := controller.NewPlayerController(playerSvc, scrimsSvc)
	discordController := controller.NewDiscordController(discordSvc)
	authController := controller.NewAuthController(authSvc)

	e.Use(middleware.Timeout())
	e.Use(middleware.Trace())
	e.Use(middleware.Logger())
	e.Use(middleware.Prepare(app.Cfg))
	e.Use(middleware.Auth(app.Cfg.JWTSecret))
	e.Use(middleware.AllowOrigin(app.Cfg))

	e.GET("/", static.Index)
	e.File("/style", "tmp/tailwind.css")
	e.File("/wasm_exec", "tmp/wasm_exec.js")
	e.File("/wasm", "tmp/app.wasm")

	discord := e.Group("/discord")
	discord.Use(middleware.Restrict(model.RoleBot))
	discord.POST("/verify", discordController.Verify)
	discord.GET("/member/:id", discordController.Member)
	discord.PATCH("/daily/:id", discordController.Daily)
	discord.DELETE("/revoke/:id", discordController.Revoke)
	discord.GET("/leaderboard/:page", discordController.Leaderboard)

	auth := e.Group("/auth")
	auth.POST("/bot", authController.Bot)

	player := e.Group("/player")
	player.GET("/exists/:name", playerController.Exists)

	scrims := player.Group("/scrims")
	scrims.GET("/:name", scrimsController.Player)
	//scrims.GET("/leaderboard/:page", scrimsController.Leaderboard)
	//scrims.GET("/scrim", )
}
