package controller

import (
	"github.com/NoBypass/fds/internal/backend/service"
	"github.com/NoBypass/fds/internal/pkg/utils"
	"github.com/labstack/echo/v4"
	"net/http"
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
	errCh := c.service.Request(ctx)

	name := ctx.Param("name")
	conn := utils.NewSSEConn(ctx.Response())

	for {
		select {
		case err := <-errCh:
			return conn.Err(err, ctx)
		//case player := <-c.service.PlayerFromDB(name):
		//	err := conn.Send(http.StatusOK, player)
		//	if err != nil {
		//		return err
		//	}
		case player := <-c.service.PlayerFromAPI(name):
			return conn.Send(http.StatusOK, player)
		}
	}
}
