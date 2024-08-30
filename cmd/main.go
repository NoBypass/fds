package main

import (
	"github.com/NoBypass/fds/internal/adapter"
	"github.com/NoBypass/fds/internal/common/env"
	"github.com/NoBypass/fds/internal/common/trace"
	"github.com/NoBypass/fds/internal/common/version"
	"github.com/NoBypass/fds/internal/port"
	"github.com/NoBypass/mincache"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetLevel(log.INFO)
	e.Logger.Print(version.Banner)

	cfg := env.Read()

	c := trace.SetupTracer(cfg)
	defer c.Close()

	cache := mincache.New()

	db := adapter.ConnectSurreal(cfg)

	go port.RunAPI(e, cfg, &port.Controllers{
		Scrims:  initScrimsController(db, cache),
		Player:  initPlayerController(db, cache),
		Discord: initDiscordController(db, cache, cfg.JwtSecret, cfg.BotPassword, cfg.HypixelAPIKey),
	})

	port.RunBot(cfg, cache)

	select {}
}
