package controller

import (
	"github.com/NoBypass/fds/internal/backend/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type PlayerController interface {
	Exists(c echo.Context) error
}

type playerController struct {
	svc       service.PlayerService
	scrimsSvc service.ScrimsService
	// TODO hypixel service
}

func NewPlayerController(svc service.PlayerService, scrimsSvc service.ScrimsService) PlayerController {
	return &playerController{
		svc:       svc,
		scrimsSvc: scrimsSvc,
	}
}

func (c playerController) Exists(ctx echo.Context) error {
	c.svc.Setup(ctx)
	c.scrimsSvc.Setup(ctx)

	name := ctx.Param("name")

	var actual string
	player, err := c.svc.FromDB(name)
	if err != nil {
		return err
	} else if player == nil || player.DisplayName == "" {
		player, err := c.scrimsSvc.PlayerFromAPI(name)
		if err != nil {
			return err
		} else if player.Data == nil {
			return echo.NewHTTPError(http.StatusNotFound, "scrims network: player not found")
		}
		actual = player.Data.Username
		go c.scrimsSvc.PersistPlayer(player)
	} else {
		actual = player.DisplayName
	}

	return ctx.String(http.StatusOK, actual)
}
