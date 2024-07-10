package port

import (
	"fmt"
	"github.com/NoBypass/fds/internal/common/env"
	"github.com/NoBypass/fds/internal/port/controller"
	"github.com/NoBypass/fds/internal/port/middleware"
	"github.com/NoBypass/mincache"
	"github.com/labstack/echo/v4"
	"time"
)

func RunAPI(e *echo.Echo, cfg *env.Env, controllers *Controllers) {
	cache := mincache.New()

	mwc := middleware.NewCacheMiddleware(cache)

	e.Use(middleware.AllowOrigin(cfg))
	e.Use(middleware.Timeout())
	e.Use(middleware.Logger())
	e.Use(middleware.Prepare(cfg))
	e.Use(middleware.Auth(cfg))
	e.Use(middleware.Trace())
	e.Use(middleware.Error())
	e.Use(middleware.Recover())

	player := e.Group("/player")
	player.GET("/exists/:name", playerController.Exists)

	scrims := player.Group("/scrims", mwc.Cache(5*time.Minute))
	scrims.GET("/:name/overview", controllers.Scrims.Player)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}

type Controllers struct {
	Scrims *controller.ScrimsController
}
