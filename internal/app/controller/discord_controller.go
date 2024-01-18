package controller

import (
	"errors"
	"fmt"
	"github.com/NoBypass/fds/internal/app/errs"
	"github.com/NoBypass/fds/internal/app/repository"
	"github.com/NoBypass/fds/internal/app/wrap"
	"github.com/NoBypass/fds/internal/pkg/conf"
	"github.com/NoBypass/fds/internal/pkg/consts"
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/surrealdb/surrealdb.go"
	"net/http"
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
	e := wrap.Error(&c)

	if err != nil {
		return e.BadRequest("invalid request query")
	}

	err = r.Create(&input)
	if err != nil {
		return err
	}
	return e.Success()
}

func (r *discordController) Daily(c echo.Context) error {
	id := c.Param("id")
	e := wrap.Error(&c)

	member, err := r.ClaimDaily(id)
	if err != nil {
		var claimedErr *errs.ClaimedError
		if errors.Is(err, surrealdb.ErrNoRow) {
			return e.NotFound(fmt.Sprintf("user with id '%s'", id))
		} else if errors.As(err, &claimedErr) {
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
	e := wrap.Error(&c)

	if err != nil {
		return e.BadRequest("invalid request query")
	}

	if input.Pwd != c.Get("config").(*conf.Config).BotPwd {
		return c.String(http.StatusForbidden, "invalid password")
	}

	claims := jwt.RegisteredClaims{
		Issuer:   consts.JWTCore,
		Subject:  consts.JWTBot,
		Audience: []string{consts.JWTBot},
		IssuedAt: jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(c.Get("config").(*conf.Config).JWTSecret))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": signedToken,
	})
}
