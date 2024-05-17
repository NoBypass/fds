package controller

import (
	"github.com/NoBypass/fds/internal/backend/service"
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthController interface {
	Bot(echo.Context) error
}

type authController struct {
	service service.AuthService
}

func NewAuthController(svc service.AuthService) AuthController {
	return &authController{
		service: svc,
	}
}

func (c authController) Bot(ctx echo.Context) error {
	errCh := c.service.Request(ctx)

	var input model.DiscordBotLoginRequest
	err := ctx.Bind(&input)
	if err != nil {
		return err
	}

	if !c.service.BotPwdIsValid(input.Pwd) {
		return ctx.JSON(http.StatusUnauthorized, "Invalid password")
	}

	claimsCh := c.service.BotClaims(input.Sub)
	token := c.service.SignJWT(claimsCh)

	select {
	case err := <-errCh:
		return err
	case token := <-token:
		return ctx.JSON(http.StatusOK, model.DiscordBotLoginResponse{
			Token: token,
		})
	}
}
