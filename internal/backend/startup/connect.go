package startup

import (
	"github.com/NoBypass/fds/internal/backend/database"
	"github.com/NoBypass/fds/internal/backend/middleware"
	"github.com/NoBypass/fds/internal/external"
	"github.com/NoBypass/fds/internal/pkg/utils"
	"github.com/NoBypass/mincache"
	"github.com/labstack/echo/v4"
	"io"
)

type App struct {
	Echo          *echo.Echo
	DB            database.Client
	HypixelClient *external.HypixelAPIClient
	Closer        io.Closer
	Cfg           *utils.Config
	Cache         *mincache.Cache
}

func Connect(e *echo.Echo) *App {
	closer := middleware.StartTracer()
	e.Logger.Info("✓ Started tracer")

	cfg := utils.ReadConfig()
	e.Logger.Info("✓ Loaded config")

	db := database.Connect(cfg)
	e.Logger.Info("✓ Connected to SurrealDB")

	cache := mincache.New()
	e.Logger.Info("✓ Started cache")

	hypixelClient := external.NewHypixelAPIClient(cache, cfg.HypixelAPIKey)
	e.Logger.Info("✓ Connected to Hypixel API")

	return &App{
		Echo:          e,
		DB:            db,
		HypixelClient: hypixelClient,
		Closer:        closer,
		Cfg:           cfg,
		Cache:         cache,
	}
}
