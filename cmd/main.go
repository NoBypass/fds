package main

import (
	"github.com/NoBypass/fds/internal/adapter"
	"github.com/NoBypass/fds/internal/common"
	"github.com/NoBypass/fds/internal/common/env"
	"github.com/NoBypass/fds/internal/port"
	"github.com/NoBypass/mincache"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetLevel(log.INFO)
	e.Logger.Print(common.Banner)

	cfg := env.Read()

	db := adapter.ConnectSurreal(cfg)

	cache := mincache.New()

	go port.RunAPI(e, cfg, &port.Controllers{
		Scrims:   initScrimsController(db, cache),
		Player:   initPlayerController(db, cache),
		Skyblock: initSkyblockController(db, cache, cfg.HypixelAPIKey),
	})

	select {}
}
