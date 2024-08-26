package main

import (
	"github.com/NoBypass/fds/internal/adapter"
	"github.com/NoBypass/fds/internal/adapter/operations"
	"github.com/NoBypass/fds/internal/app"
	"github.com/NoBypass/fds/internal/port/controller"
	"github.com/NoBypass/mincache"
)

func initPlayerController(db adapter.Database, cache *mincache.Cache) *controller.PlayerController {
	scrimsRepository := operations.NewScrimsRepository(db)
	scrimsAPI := adapter.NewScrimsAPI(cache)
	mojangAPI := adapter.NewMojangAPI(cache)
	minecraftUseCase := app.NewMinecraftUseCase(scrimsRepository, scrimsAPI, mojangAPI)
	playerController := controller.NewPlayerController(minecraftUseCase)
	return playerController
}

func initScrimsController(db adapter.Database, cache *mincache.Cache) *controller.ScrimsController {
	scrimsRepository := operations.NewScrimsRepository(db)
	scrimsAPI := adapter.NewScrimsAPI(cache)
	mojangAPI := adapter.NewMojangAPI(cache)
	minecraftUseCase := app.NewMinecraftUseCase(scrimsRepository, scrimsAPI, mojangAPI)
	scrimsController := controller.NewScrimsController(minecraftUseCase)
	return scrimsController
}

func initDiscordController(db adapter.Database, cache *mincache.Cache, jwtSecret, pwd, hypixelKey string) *controller.DiscordController {
	discordRepository := operations.NewDiscordRepository(db)
	hypixelAPI := adapter.NewHypixelAPI(cache, hypixelKey)
	discordUseCase := app.NewDiscordUseCase(discordRepository, hypixelAPI, jwtSecret, pwd)
	discordController := controller.NewDiscordController(discordUseCase)
	return discordController
}
