package service

import (
	"encoding/json"
	"github.com/NoBypass/fds/internal/app/errs"
	"github.com/NoBypass/fds/internal/app/repository"
	"github.com/NoBypass/fds/internal/pkg/conf"
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/surrealdb/surrealdb.go"
	"math"
	"math/rand"
	"net/http"
	"time"
)

type DiscordService interface {
	Service
	ParseBotLogin(c echo.Context) <-chan model.DiscordBotLoginInput
	ParseVerify(c echo.Context) <-chan model.DiscordVerifyInput
	ParseDaily(c echo.Context) <-chan string

	Persist(profileCh <-chan model.MojangProfile, memberCh <-chan model.DiscordMember, playerCh <-chan model.HypixelPlayer) <-chan struct{}

	VerifyHypixelSocials(member <-chan model.DiscordMember, player <-chan model.HypixelPlayerResponse) (<-chan model.DiscordMember, <-chan model.HypixelPlayer)
	FetchHypixelPlayer(inputCh <-chan model.MojangProfile) (<-chan model.HypixelPlayerResponse, <-chan model.MojangProfile)
	FetchMojangProfile(inputCh <-chan model.DiscordVerifyInput) (<-chan model.MojangProfile, <-chan model.DiscordMember)
	GiveXP(member <-chan model.DiscordMember, xp <-chan float64) <-chan model.DiscordMember
	GetJWT(input <-chan model.DiscordBotLoginInput) <-chan string
	CheckDaily(member <-chan model.DiscordMember) <-chan float64
	GetMember(id <-chan string) <-chan model.DiscordMember
}

type discordService struct {
	service
	repo        repository.DiscordRepository
	mojangRepo  repository.MojangRepository
	hypixelRepo repository.HypixelRepository
	config      *conf.Config
}

func NewDiscordService(db *surrealdb.DB, config *conf.Config) DiscordService {
	return &discordService{
		hypixelRepo: repository.NewHypixelRepository(db),
		repo:        repository.NewDiscordRepository(db),
		mojangRepo:  repository.NewMojangRepository(db),
		config:      config,
	}
}

func (s *discordService) InjectErrorChan() <-chan error {
	ch := make(chan error)
	s.errCh = ch
	return ch
}

func (s *discordService) ParseVerify(c echo.Context) <-chan model.DiscordVerifyInput {
	inputCh := make(chan model.DiscordVerifyInput)

	go func() {
		defer close(inputCh)

		var input model.DiscordVerifyInput
		err := c.Bind(&input)
		if err != nil {
			s.errCh <- errs.BadRequest("error parsing input")
			return
		}

		inputCh <- input
	}()

	return inputCh
}

func (s *discordService) ParseBotLogin(c echo.Context) <-chan model.DiscordBotLoginInput {
	inputCh := make(chan model.DiscordBotLoginInput)

	go func() {
		defer close(inputCh)

		var input model.DiscordBotLoginInput
		err := c.Bind(&input)
		if err != nil {
			s.errCh <- errs.BadRequest("error parsing input")
			return
		}

		inputCh <- input
	}()

	return inputCh
}

func (s *discordService) ParseDaily(c echo.Context) <-chan string {
	idCh := make(chan string)

	go func() {
		defer close(idCh)

		id := c.Param("id")
		if id == "" {
			s.errCh <- errs.BadRequest("error parsing input")
			return
		}
		idCh <- id
	}()

	return idCh
}

func (s *discordService) GetMember(id <-chan string) <-chan model.DiscordMember {
	memberCh := make(chan model.DiscordMember)

	go func() {
		defer close(memberCh)

		member, err := s.repo.Get(<-id)
		if err != nil {
			s.errCh <- err
			return
		}

		memberCh <- *member
	}()

	return memberCh
}

func (s *discordService) CheckDaily(memberCh <-chan model.DiscordMember) <-chan float64 {
	xpCh := make(chan float64)

	go func() {
		defer close(xpCh)

		m := <-memberCh
		if m.CanClaimDaily() {
			m.AddXP(math.Round(rand.Float64() * 500.0 * (1.0 + float64(m.Streak)*0.1)))
			m.LastDailyAt = time.Now().UnixMilli()
			m.Streak++
		} else {
			s.errCh <- errs.TooManyRequests("user has already claimed their daily reward")
		}
	}()

	return xpCh
}

func (s *discordService) GiveXP(memberCh <-chan model.DiscordMember, xp <-chan float64) <-chan model.DiscordMember {
	out := make(chan model.DiscordMember)

	go func() {
		defer close(out)

		member := <-memberCh
		member.AddXP(<-xp)
		err := s.repo.Update(member.DiscordID, &member)
		if err != nil {
			s.errCh <- err
			return
		}
		out <- member
	}()

	return memberCh
}

func (s *discordService) GetJWT(input <-chan model.DiscordBotLoginInput) <-chan string {
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
				s.errCh <- err
				return
			}
			tokenCh <- signedToken
		} else {
			s.errCh <- errs.Unauthorized("invalid password")
			return
		}
	}()

	return tokenCh
}

func (s *discordService) FetchMojangProfile(inputCh <-chan model.DiscordVerifyInput) (<-chan model.MojangProfile, <-chan model.DiscordMember) {
	profileCh := make(chan model.MojangProfile)
	memberCh := make(chan model.DiscordMember)

	go func() {
		defer close(profileCh)
		defer close(memberCh)

		input := <-inputCh
		resp, err := http.Get("https://api.mojang.com/users/profiles/minecraft/" + input.Nick)
		if err != nil {
			s.errCh <- err
			return
		}

		var profile model.MojangProfile
		err = json.NewDecoder(resp.Body).Decode(&profile)
		if err != nil {
			s.errCh <- err
			return
		}

		profileCh <- profile
		memberCh <- model.DiscordMember{
			DiscordID: input.ID,
			Name:      input.Name,
			Nick:      profile.Name,
		}
	}()

	return profileCh, memberCh
}

func (s *discordService) FetchHypixelPlayer(inputCh <-chan model.MojangProfile) (<-chan model.HypixelPlayerResponse, <-chan model.MojangProfile) {
	playerCh := make(chan model.HypixelPlayerResponse)
	profileCh := make(chan model.MojangProfile)

	go func() {
		defer close(playerCh)

		input := <-inputCh
		profileCh <- input
		req, err := http.NewRequest(http.MethodGet, "https://api.hypixel.net/player?uuid="+input.UUID, nil)
		if err != nil {
			s.errCh <- err
			return
		}

		req.Header.Add("API-Key", s.config.HypixelAPIKey)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			s.errCh <- err
			return
		}

		var player model.HypixelPlayerResponse
		err = json.NewDecoder(resp.Body).Decode(&player)
		if err != nil {
			s.errCh <- err
			return
		}

		playerCh <- player
	}()

	return playerCh, profileCh
}

func (s *discordService) VerifyHypixelSocials(memberCh <-chan model.DiscordMember, playerCh <-chan model.HypixelPlayerResponse) (<-chan model.DiscordMember, <-chan model.HypixelPlayer) {
	outMemberCh := make(chan model.DiscordMember)
	outPlayerCh := make(chan model.HypixelPlayer)

	go func() {
		defer close(outPlayerCh)
		defer close(outMemberCh)

		member := <-memberCh
		player := <-playerCh

		if player.Success {
			if player.Player.SocialMedia.Links.Discord == member.Name {
				outMemberCh <- member
				outPlayerCh <- model.HypixelPlayer{
					Date: time.Now().Format(time.RFC3339),
					UUID: player.Player.UUID,
				}
			} else {
				s.errCh <- errs.Forbidden("discord id does not match hypixel socials")
				return
			}
		} else {
			s.errCh <- errs.NotFound("hypixel: player not found")
			return
		}
	}()

	return outMemberCh, outPlayerCh
}

func (s *discordService) Persist(profileCh <-chan model.MojangProfile, memberCh <-chan model.DiscordMember, playerCh <-chan model.HypixelPlayer) <-chan struct{} {
	done := make(chan struct{})

	go func() {
		defer close(done)

		var (
			p model.MojangProfile
			m model.DiscordMember
			h model.HypixelPlayer
		)
		for i := 0; i < 3; i++ {
			select {
			case profile := <-profileCh:
				err := s.mojangRepo.Create(&profile)
				if err != nil {
					s.errCh <- err
					return
				}
				p = profile
			case member := <-memberCh:
				err := s.repo.Create(&member)
				if err != nil {
					s.errCh <- err
					return
				}
				m = member
			case player := <-playerCh:
				err := s.hypixelRepo.Create(&player)
				if err != nil {
					s.errCh <- err
					return
				}
				h = player
			}
		}

		err := s.repo.RelatePlayedWith(&p, &h)
		if err != nil {
			s.errCh <- err
			return
		}

		err = s.repo.RelateVerifiedWith(&m, &h)
		if err != nil {
			s.errCh <- err
			return
		}

		done <- struct{}{}
	}()

	return done
}
