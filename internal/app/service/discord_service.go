package service

import (
	"github.com/NoBypass/fds/internal/app/errs"
	"github.com/NoBypass/fds/internal/app/repository"
	"github.com/NoBypass/fds/internal/pkg/conf"
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/surrealdb/surrealdb.go"
	"math"
	"math/rand"
	"time"
)

type DiscordService interface {
	ParseVerify(c echo.Context, err chan<- error) <-chan model.DiscordVerifyInput
	ParseDaily(c echo.Context, err chan<- error) <-chan string
	ParseBotLogin(c echo.Context, err chan<- error) <-chan model.DiscordBotLoginInput

	CreateMember(member <-chan model.DiscordVerifyInput, err chan<- error) <-chan model.DiscordMember
	GetMember(id <-chan string, err chan<- error) <-chan model.DiscordMember

	GiveXP(member <-chan model.DiscordMember, xp <-chan float64, err chan<- error) <-chan model.DiscordMember
	CheckDaily(member <-chan model.DiscordMember, err chan<- error) <-chan float64
	GetJWT(input <-chan model.DiscordBotLoginInput, config *conf.Config, err chan<- error) <-chan string
}

type discordService struct {
	repo repository.DiscordRepository
}

func NewDiscordService(db *surrealdb.DB) DiscordService {
	return &discordService{
		repository.NewDiscordRepository(db),
	}
}

func (s *discordService) ParseVerify(c echo.Context, errCh chan<- error) <-chan model.DiscordVerifyInput {
	inputCh := make(chan model.DiscordVerifyInput)

	go func() {
		defer close(inputCh)

		var input model.DiscordVerifyInput
		err := c.Bind(&input)
		if err != nil {
			errCh <- errs.BadRequest("error parsing input")
			return
		}

		inputCh <- input
	}()

	return inputCh
}

func (s *discordService) ParseBotLogin(c echo.Context, errCh chan<- error) <-chan model.DiscordBotLoginInput {
	inputCh := make(chan model.DiscordBotLoginInput)

	go func() {
		defer close(inputCh)

		var input model.DiscordBotLoginInput
		err := c.Bind(&input)
		if err != nil {
			errCh <- errs.BadRequest("error parsing input")
			return
		}

		inputCh <- input
	}()

	return inputCh
}

func (s *discordService) CreateMember(member <-chan model.DiscordVerifyInput, errCh chan<- error) <-chan model.DiscordMember {
	memberCh := make(chan model.DiscordMember)

	go func() {
		defer close(memberCh)

		input := <-member
		discordMember := model.DiscordMember{
			Nick: input.Nick,
		}
		err := s.repo.Create(&discordMember)
		if err != nil {
			errCh <- errs.BadRequest("error parsing input")
			return
		}

		memberCh <- discordMember
	}()

	return memberCh
}

func (s *discordService) ParseDaily(c echo.Context, errCh chan<- error) <-chan string {
	idCh := make(chan string)

	go func() {
		defer close(idCh)

		id := c.Param("id")
		if id == "" {
			errCh <- errs.BadRequest("error parsing input")
			return
		}
		idCh <- id
	}()

	return idCh
}

func (s *discordService) GetMember(id <-chan string, errCh chan<- error) <-chan model.DiscordMember {
	memberCh := make(chan model.DiscordMember)

	go func() {
		defer close(memberCh)

		member, err := s.repo.Get(<-id)
		if err != nil {
			errCh <- err
			return
		}

		memberCh <- *member
	}()

	return memberCh
}

func (s *discordService) CheckDaily(member <-chan model.DiscordMember, errCh chan<- error) <-chan float64 {
	xpCh := make(chan float64)

	go func() {
		defer close(xpCh)

		m := <-member
		if m.CanClaimDaily() {
			m.AddXP(math.Round(rand.Float64() * 500.0 * (1.0 + float64(m.Streak)*0.1)))
			m.LastDailyClaim = time.Now().UnixMilli()
			m.Streak++
		} else {
			errCh <- errs.TooManyRequests("user has already claimed their daily reward")
		}
	}()

	return xpCh
}

func (s *discordService) GiveXP(member <-chan model.DiscordMember, xp <-chan float64, errCh chan<- error) <-chan model.DiscordMember {
	memberCh := make(chan model.DiscordMember)

	go func() {
		defer close(memberCh)

		m := <-member
		m.AddXP(<-xp)
		err := s.repo.Update(&m)
		if err != nil {
			errCh <- err
			return
		}
		memberCh <- m
	}()

	return memberCh
}

func (s *discordService) GetJWT(input <-chan model.DiscordBotLoginInput, config *conf.Config, errCh chan<- error) <-chan string {
	tokenCh := make(chan string)

	go func() {
		defer close(tokenCh)

		input := <-input
		if input.Pwd == config.BotPwd {
			claims := jwt.RegisteredClaims{
				Issuer:   "fds",
				Subject:  "bot",
				Audience: []string{"bot"},
				IssuedAt: jwt.NewNumericDate(time.Now()),
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			signedToken, err := token.SignedString([]byte(config.JWTSecret))
			if err != nil {
				errCh <- err
				return
			}
			tokenCh <- signedToken
		} else {
			errCh <- errs.Unauthorized("invalid password")
			return
		}
	}()

	return tokenCh
}
