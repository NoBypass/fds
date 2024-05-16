package controller

import (
	"github.com/NoBypass/fds/internal/backend/service"
	"github.com/NoBypass/fds/internal/pkg/utils"
	"github.com/labstack/echo/v4"
)

type ScrimsController interface {
	Leaderboard(c echo.Context) error
	Player(c echo.Context) error
}

type scrimsController struct {
	service service.ScrimsService
}

func NewScrimsController(svc service.ScrimsService) ScrimsController {
	return &scrimsController{
		service: svc,
	}
}

func (c scrimsController) Leaderboard(ctx echo.Context) error {
	// TODO
	return nil
}

func (c scrimsController) Player(ctx echo.Context) error {
	name := ctx.Param("name")
	conn := utils.NewSSEConn(ctx.Response())

	for {
		select {
		case err := <-c.service.Error():
			return err
		case player := <-c.service.PlayerFromDB(name):
			err := conn.Send(player)
			if err != nil {
				return err
			}
		case player := <-c.service.PlayerFromAPI(name):
			return conn.Send(player)
		}
	}
}
