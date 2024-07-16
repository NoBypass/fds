//go:build wireinject
// +build wireinject

package main

import (
	"github.com/NoBypass/fds/internal/adapter"
	"github.com/NoBypass/fds/internal/adapter/operations"
	"github.com/NoBypass/fds/internal/app"
	"github.com/NoBypass/fds/internal/port/controller"
	"github.com/NoBypass/mincache"
	"github.com/google/wire"
)

func initScrimsController(db adapter.Database, cache *mincache.Cache) *controller.ScrimsController {
	wire.Build(
		wire.Bind(new(app.MinecraftRepository), new(*operations.ScrimsRepository)),
		wire.Bind(new(app.ScrimsService), new(*adapter.ScrimsAPI)),
		wire.Bind(new(app.MojangService), new(*adapter.MojangAPI)),

		controller.NewScrimsController,
		app.NewMinecraftUseCase,
		operations.NewScrimsRepository,
		adapter.NewMojangAPI,
		adapter.NewScrimsAPI)
	return nil
}

func initPlayerController(db adapter.Database, cache *mincache.Cache) *controller.PlayerController {
	wire.Build(
		wire.Bind(new(app.MinecraftRepository), new(*operations.ScrimsRepository)),
		wire.Bind(new(app.ScrimsService), new(*adapter.ScrimsAPI)),
		wire.Bind(new(app.MojangService), new(*adapter.MojangAPI)),

		controller.NewPlayerController,
		app.NewMinecraftUseCase,
		operations.NewScrimsRepository,
		adapter.NewMojangAPI,
		adapter.NewScrimsAPI)
	return nil
}
