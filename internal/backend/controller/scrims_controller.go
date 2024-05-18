package controller

import (
	"github.com/NoBypass/fds/internal/backend/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ScrimsController interface {
	Leaderboard(c echo.Context) error
	Player(c echo.Context) error
}

type scrimsController struct {
	service   service.ScrimsService
	mojangSvc service.MojangService
}

func NewScrimsController(svc service.ScrimsService, mojangSvc service.MojangService) ScrimsController {
	return &scrimsController{
		mojangSvc: mojangSvc,
		service:   svc,
	}
}

func (c scrimsController) Leaderboard(ctx echo.Context) error {
	// TODO
	return nil
}

func (c scrimsController) Player(ctx echo.Context) error {
	c.service.Setup(ctx)
	c.mojangSvc.Setup(ctx)

	name := ctx.Param("name")

	rawPlayer, err := c.service.PlayerFromAPI(name)
	if err != nil {
		return err
	} else if rawPlayer.Data == nil {
		return echo.NewHTTPError(http.StatusNotFound, "scrims network: player not found")
	}

	dbPlayer, err := c.mojangSvc.PlayerFromDB(name, "scrims_data.date", "uuid")
	if err != nil {
		return err
	} else if dbPlayer == nil || dbPlayer.UUID == "" {
		player, err := c.mojangSvc.PlayerFromAPI(name)
		if err != nil {
			return err
		}

		dbPlayer, err = c.mojangSvc.PersistPlayer(player)
		if err != nil {
			return err
		}
	}

	player, err := c.service.PersistScrimsPlayer(rawPlayer, dbPlayer)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, player)
}
