package service

import (
	"encoding/base64"
	"encoding/json"
	"github.com/NoBypass/fds/internal/app/errs"
	"github.com/NoBypass/fds/internal/app/repository"
	"github.com/NoBypass/fds/internal/pkg/conf"
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/surrealdb/surrealdb.go"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type DiscordService interface {
	ParseBotLogin(c echo.Context, err chan<- error) <-chan model.DiscordBotLoginInput
	ParseVerify(c echo.Context, err chan<- error) <-chan model.DiscordVerifyInput
	ParseDaily(c echo.Context, err chan<- error) <-chan string

	CreateMemberAndProfile(profile <-chan model.MojangProfile, member <-chan model.DiscordMember, err chan<- error) <-chan struct{}
	GetMember(id <-chan string, err chan<- error) <-chan model.DiscordMember

	FetchMojangProfile(inputCh <-chan model.DiscordVerifyInput, err chan<- error) (<-chan model.MojangProfile, <-chan model.DiscordMember)
	FetchHypixelPlayer(inputCh <-chan model.MojangProfile, err chan<- error) (<-chan model.HypixelPlayer, <-chan model.DiscordMember)
	GiveXP(member <-chan model.DiscordMember, xp <-chan float64, err chan<- error) <-chan model.DiscordMember
	GetJWT(input <-chan model.DiscordBotLoginInput, err chan<- error) <-chan string
	CheckDaily(member <-chan model.DiscordMember, err chan<- error) <-chan float64
}

type discordService struct {
	repo       repository.DiscordRepository
	mojangRepo repository.MojangRepository
	config     *conf.Config
}

func NewDiscordService(db *surrealdb.DB, config *conf.Config) DiscordService {
	return &discordService{
		repository.NewDiscordRepository(db),
		repository.NewMojangRepository(db),
		config,
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

func (s *discordService) CreateMemberAndProfile(profileCh <-chan model.MojangProfile, memberCh <-chan model.DiscordMember, errCh chan<- error) <-chan struct{} {
	done := make(chan struct{})

	go func() {
		defer close(done)

		profile := <-profileCh
		err := s.mojangRepo.Create(&profile)
		if err != nil && !strings.Contains(err.Error(), "already exists") {
			errCh <- err
			return
		}

		member := <-memberCh
		err = s.repo.Create(&member, &profile)
		if err != nil {
			errCh <- err
			return
		}
	}()

	return done
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

func (s *discordService) CheckDaily(memberCh <-chan model.DiscordMember, errCh chan<- error) <-chan float64 {
	xpCh := make(chan float64)

	go func() {
		defer close(xpCh)

		m := <-memberCh
		if m.CanClaimDaily() {
			m.AddXP(math.Round(rand.Float64() * 500.0 * (1.0 + float64(m.Streak)*0.1)))
			m.LastDailyAt = time.Now().UnixMilli()
			m.Streak++
		} else {
			errCh <- errs.TooManyRequests("user has already claimed their daily reward")
		}
	}()

	return xpCh
}

func (s *discordService) GiveXP(memberCh <-chan model.DiscordMember, xp <-chan float64, errCh chan<- error) <-chan model.DiscordMember {
	out := make(chan model.DiscordMember)

	go func() {
		defer close(out)

		m := <-memberCh
		m.AddXP(<-xp)
		err := s.repo.Update(&m)
		if err != nil {
			errCh <- err
			return
		}
		out <- m
	}()

	return memberCh
}

func (s *discordService) GetJWT(input <-chan model.DiscordBotLoginInput, errCh chan<- error) <-chan string {
	tokenCh := make(chan string)

	go func() {
		defer close(tokenCh)

		input := <-input
		if input.Pwd == s.config.BotPwd {
			claims := jwt.RegisteredClaims{
				Issuer:   "fds",
				Subject:  "bot",
				Audience: []string{"bot"},
				IssuedAt: jwt.NewNumericDate(time.Now()),
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			signedToken, err := token.SignedString([]byte(s.config.JWTSecret))
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

func (s *discordService) FetchMojangProfile(inputCh <-chan model.DiscordVerifyInput, errCh chan<- error) (<-chan model.MojangProfile, <-chan model.DiscordMember) {
	profileCh := make(chan model.MojangProfile)
	memberCh := make(chan model.DiscordMember)

	go func() {
		defer close(profileCh)
		defer close(memberCh)

		input := <-inputCh
		resp, err := http.Get("https://api.mojang.com/users/profiles/minecraft/" + input.Nick)
		if err != nil {
			errCh <- err
			return
		}

		var profile model.MojangProfile
		err = json.NewDecoder(resp.Body).Decode(&profile)
		if err != nil {
			errCh <- err
			return
		}

		profileCh <- profile
		memberCh <- model.DiscordMember{
			ID:   input.ID,
			Nick: profile.Name,
		}
	}()

	return profileCh, memberCh
}

func (s *discordService) FetchHypixelPlayer(inputCh <-chan model.MojangProfile, errCh chan<- error) (<-chan model.HypixelPlayer, <-chan model.DiscordMember) {
	playerCh := make(chan model.HypixelPlayer)
	memberCh := make(chan model.DiscordMember)

	go func() {
		defer close(playerCh)
		defer close(memberCh)

		input := <-inputCh
		resp, err := http.Get("https://api.hypixel.net/player?key=" + s.config.HypixelAPIKey + "&uuid=" + input.ID)
		if err != nil {
			errCh <- err
			return
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			errCh <- err
			return
		}

		base64Body := base64.StdEncoding.EncodeToString(body)
		var player model.HypixelPlayerResponse
		err = json.Unmarshal(body, &player)
		if err != nil {
			errCh <- err
			return
		}

		var respBody interface{}
		err = json.NewDecoder(resp.Body).Decode(&respBody)
		if err != nil {
			errCh <- err
			return
		}

		playerCh <- model.HypixelPlayer(base64Body)
		memberCh <- model.DiscordMember{
			ID:   input.ID,
			Nick: player.Player.DisplayName,
		}
	}()

	return playerCh, memberCh
}
