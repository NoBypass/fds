package controller

import (
	"github.com/NoBypass/fds/internal/app"
	"github.com/labstack/echo/v4"
	"net/http"
)

type SkyblockController struct {
	uc *app.SkyblockUseCase
}

func NewSkyblockController(uc *app.SkyblockUseCase) *SkyblockController {
	return &SkyblockController{
		uc: uc,
	}
}

func (ct *SkyblockController) AuctionsOverview(ctx echo.Context) error {
	auctions, err := ct.uc.GetAuctions(ctx.Request().Context())
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, auctions)
}
