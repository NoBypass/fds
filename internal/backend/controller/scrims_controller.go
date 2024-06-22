package controller

import (
	"github.com/NoBypass/fds/internal/backend/service"
	"github.com/NoBypass/fds/internal/backend/tracing"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ScrimsController interface {
	Leaderboard(c echo.Context) error
	Player(c echo.Context) error
	Overview(c echo.Context) error
}

type scrimsController struct {
	tracing.Tracable

	service      service.ScrimsService
	mojangSvc    service.MojangService
	minecraftSvc service.MinecraftService
}

func NewScrimsController(svc service.ScrimsService, mojangSvc service.MojangService, minecraftSvc service.MinecraftService) ScrimsController {
	return &scrimsController{
		mojangSvc:    mojangSvc,
		service:      svc,
		minecraftSvc: minecraftSvc,
		Tracable:     tracing.NewTracable(),
	}
}

func (ct scrimsController) Overview(c echo.Context) error {
	ctx := c.Request().Context()

	name := c.Param("name")

	player, err := ct.minecraftSvc.RemoteScrimsStats(ctx, name)
	if err != nil {
		return err
	}

	totals, err := ct.minecraftSvc.TotalScrimsStats(ctx, player)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]any{
		"player": player,
		"totals": totals,
	})
}

func (ct scrimsController) Leaderboard(ctx echo.Context) error {
	// TODO
	return nil
}

func (ct scrimsController) Player(ctx echo.Context) error {
	ct.mojangSvc.Setup(ctx)

	name := ctx.Param("name")

	rawPlayer, err := ct.service.PlayerFromAPI(name)
	if err != nil {
		return err
	} else if rawPlayer.Data == nil {
		return echo.NewHTTPError(http.StatusNotFound, "scrims network: playerTimes not found")
	}

	dbPlayer, err := ct.mojangSvc.PlayerFromDB(name, "scrims_data.date", "uuid")
	if err != nil {
		return err
	} else if dbPlayer == nil || dbPlayer.UUID == "" {
		dbPlayer, err = ct.service.PersistPlayer(rawPlayer)
		if err != nil {
			return err
		}
	}

	_, err = ct.service.PersistScrimsPlayer(rawPlayer, dbPlayer)
	if err != nil {
		return err
	}

	playerTimes, err := ct.service.AllPlayerTimes(name)
	if err != nil {
		return err
	}

	player, err := ct.service.PlayerByDate(name, playerTimes[0].Date)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]any{
		"player": player.Data,
		"times":  playerTimes,
	})
}
