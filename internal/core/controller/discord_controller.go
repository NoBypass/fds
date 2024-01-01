package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/surrealdb/surrealdb.go"
	"net/http"
	"server/internal/core/custom_err"
	"server/internal/core/repository"
	"server/internal/pkg/model"
)

type DiscordController interface {
	Signup(c echo.Context) error
	Daily(c echo.Context) error
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
