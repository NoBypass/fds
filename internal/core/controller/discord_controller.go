package controller

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/surrealdb/surrealdb.go"
	"net/http"
	"server/internal/core/custom_err"
	"server/internal/core/repository"
	"server/internal/pkg/conf"
	"server/internal/pkg/consts"
	"server/internal/pkg/model"
	"time"
)

type DiscordController interface {
	Signup(c echo.Context) error
	Daily(c echo.Context) error
	BotLogin(c echo.Context) error
}

type discordController struct {
	repository.DiscordRepository
}

func NewDiscordController(db *surrealdb.DB) DiscordController {
	return &discordController{
		repository.NewDiscordRepository(db),
	}
}

func (r *discordController) Signup(c echo.Context) error {
	var input model.DiscordSignupInput
	err := c.Bind(&input)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid request query")
	}

	err = r.Create(&input)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "success")
}

func (r *discordController) Daily(c echo.Context) error {
	id := c.Param("id")

	member, err := r.ClaimDaily(id)
	if err != nil {
		var claimedErr *custom_err.ClaimedError
		if errors.As(err, &claimedErr) {
			return c.JSON(http.StatusForbidden, claimedErr)
		} else {
			return err
		}
	}

	return c.JSON(http.StatusOK, *member)
}

func (r *discordController) BotLogin(c echo.Context) error {
	input := model.DiscordBotLoginInput{}
	err := c.Bind(&input)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid request query")
	}

	if input.Pwd != c.Get("config").(*conf.Config).Authentication.Bot.Password {
		return c.String(http.StatusForbidden, "invalid password")
	}

	claims := jwt.RegisteredClaims{
		Issuer:   consts.JWTCore,
		Subject:  consts.JWTBot,
		Audience: []string{consts.JWTBot},
		IssuedAt: jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(c.Get("config").(*conf.Config).Authentication.Jwt.Secret))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": signedToken,
	})
}
