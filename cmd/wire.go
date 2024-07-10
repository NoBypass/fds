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
		wire.Bind(new(app.MinecraftService), new(*adapter.ScrimsAPI)),
		controller.NewScrimsController,
		app.NewMinecraftUseCase,
		operations.NewScrimsRepository,
		adapter.NewScrimsAPI)
	return nil
}
