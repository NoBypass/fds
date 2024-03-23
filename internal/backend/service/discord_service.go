package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/NoBypass/fds/internal/backend/repository"
	"github.com/NoBypass/fds/internal/pkg/conf"
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/NoBypass/fds/pkg/api"
	"github.com/NoBypass/surgo"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/surrealdb/surrealdb.go"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type DiscordService interface {
	Service
	PersistVerify(profileCh <-chan model.MojangProfile, memberCh <-chan model.DiscordMember, playerCh <-chan model.HypixelPlayer) <-chan string

	CheckIfAlreadyVerified(input *api.DiscordVerifyRequest) <-chan *api.DiscordVerifyRequest
	VerifyHypixelSocials(member <-chan model.DiscordMember, player <-chan model.HypixelPlayerResponse) (<-chan model.DiscordMember, <-chan model.HypixelPlayer)
	FetchHypixelPlayer(inputCh <-chan model.MojangProfile) (<-chan model.HypixelPlayerResponse, <-chan model.MojangProfile)
	FetchMojangProfile(inputCh <-chan *api.DiscordVerifyRequest) (<-chan model.MojangProfile, <-chan model.DiscordMember)
	GiveDaily(member <-chan model.DiscordMember) <-chan model.DiscordMember
	GetJWT(input *api.DiscordBotLoginRequest) <-chan string
	GetMember(id string) <-chan model.DiscordMember
	StrToInt(input string) <-chan int
	GetLeaderboard(page <-chan int) <-chan api.DiscordLeaderboardResponse
	Revoke(id string) <-chan *api.DiscordMemberResponse
}

type discordService struct {
	service
	config *conf.Config
}

func NewDiscordService(config *conf.Config) DiscordService {
	return &discordService{
		config: config,
	}
}

func (s *discordService) GetMember(id string) <-chan model.DiscordMember {
	memberCh := make(chan model.DiscordMember)

	go func() {
		defer close(memberCh)

		member := new(model.DiscordMember)
		err := repository.DB.Scan(member, "SELECT * FROM ONLY discord_member:$", surgo.ID{id})
		if err != nil {
			s.errCh <- err
			return
		}

		memberCh <- *member
	}()

	return memberCh
}

func (s *discordService) GiveDaily(memberCh <-chan model.DiscordMember) <-chan model.DiscordMember {
	out := make(chan model.DiscordMember)

	go func() {
		defer close(out)

		member := <-memberCh

		member.AddXP(math.Round(rand.Float64() * 500.0 * (1.0 + float64(member.Streak)*0.1)))
		member.LastDailyAt = time.Now().Add(-time.Hour * 24).Format(time.RFC3339)
		member.Streak++

		var newMember model.DiscordMember
		err := repository.DB.Scan(&newMember, "UPDATE discord_member:$ MERGE {"+
			"last_daily_at: $last_daily_at,"+
			"xp: $xp,"+
			"streak: $streak"+
			"}", member, surgo.ID{member.DiscordID})
		if err != nil {
			s.errCh <- err
			return
		}
		out <- newMember
	}()

	return memberCh
}

func (s *discordService) GetJWT(input *api.DiscordBotLoginRequest) <-chan string {
	tokenCh := make(chan string)

	go func() {
		defer close(tokenCh)

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
			s.errCh <- echo.NewHTTPError(http.StatusUnauthorized, "invalid password")
			return
		}
	}()

	return tokenCh
}

func (s *discordService) FetchMojangProfile(inputCh <-chan *api.DiscordVerifyRequest) (<-chan model.MojangProfile, <-chan model.DiscordMember) {
	profileCh := make(chan model.MojangProfile)
	memberCh := make(chan model.DiscordMember)

	go func() {
		defer close(profileCh)
		defer close(memberCh)

		input, ok := <-inputCh
		if !ok {
			return
		}
		resp, err := http.Get("https://api.mojang.com/users/profiles/minecraft/" + input.Nick)
		if err != nil {
			s.errCh <- err
			return
		}

		var profile model.MojangProfile
		err = json.NewDecoder(resp.Body).Decode(&profile)
		if err != nil {
			s.errCh <- fmt.Errorf("mojang: %s", err)
			return
		}

		profile.Date = time.Now().Format(time.RFC3339)
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

		input, ok := <-inputCh
		if !ok {
			return
		}
		profileCh <- input
		req, err := http.NewRequest(http.MethodGet, "https://api.hypixel.net/player?uuid="+input.UUID, nil)
		if err != nil {
			s.errCh <- err
			return
		}

		req.Header.Add("API-Key", s.config.HypixelAPIKey)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			s.errCh <- fmt.Errorf("hypixel: %s", err)
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

		member, ok := <-memberCh
		if !ok {
			return
		}
		player, ok := <-playerCh
		if !ok {
			return
		}

		if player.Success {
			if player.Player.SocialMedia.Links.Discord == member.Name {
				outMemberCh <- member
				outPlayerCh <- model.HypixelPlayer{
					UUID: player.Player.UUID,
					Date: time.Now().Format(time.RFC3339),
				}
			} else {
				s.errCh <- echo.NewHTTPError(http.StatusForbidden, "discord id does not match hypixel socials")
				return
			}
		} else {
			s.errCh <- echo.NewHTTPError(http.StatusNotFound, "hypixel: player not found")
			return
		}
	}()

	return outMemberCh, outPlayerCh
}

func (s *discordService) PersistVerify(profileCh <-chan model.MojangProfile, memberCh <-chan model.DiscordMember, playerCh <-chan model.HypixelPlayer) <-chan string {
	actual := make(chan string)

	go func() {
		defer close(actual)

		var (
			p model.MojangProfile
			m model.DiscordMember
			h model.HypixelPlayer
		)
		for i := 0; i < 3; i++ {
			select {
			case profile, ok := <-profileCh:
				if !ok {
					return
				}
				_, err := repository.DB.Exec("CREATE mojang_profile:$ CONTENT {"+
					"date: $date,"+
					"uuid: $uuid,"+
					"name: $name"+
					"}", profile, surgo.ID{profile.UUID, profile.Date})
				if err != nil {
					s.errCh <- err
					return
				}
				p = profile
			case member, ok := <-memberCh:
				if !ok {
					return
				}
				err := repository.DB.Scan(&m, "CREATE discord_member:$ CONTENT {"+
					"discord_id: $discord_id,"+
					"name: $name,"+
					"nick: $nick,"+
					"xp: 0,"+
					"last_daily_at: $last_daily_at,"+
					"level: 0,"+
					"streak: 0"+
					"}", member, surgo.ID{member.DiscordID, time.Now().Format(time.RFC3339)})
				if err != nil {
					s.errCh <- err
					return
				}
			case player, ok := <-playerCh:
				if !ok {
					return
				}
				err := repository.DB.Scan(&h, "CREATE hypixel_player:$ CONTENT {"+
					"uuid: $uuid,"+
					"date: $date"+
					"}", player, surgo.ID{p.UUID, p.Date})
				if err != nil {
					s.errCh <- err
					return
				}
			}
		}

		_, err := repository.DB.Exec("RELATE discord_member:$->PLAYED_WITH->mojang_account:$", surgo.ID{m.DiscordID}, surgo.ID{p.UUID})
		if err != nil {
			s.errCh <- err
			return
		}

		_, err = repository.DB.Exec("RELATE mojang_profile:$->VERIFIED_WITH->discord_member:$", surgo.ID{p.UUID}, surgo.ID{m.DiscordID})
		if err != nil {
			s.errCh <- err
			return
		}

		actual <- p.Name
	}()

	return actual
}

func (s *discordService) CheckIfAlreadyVerified(input *api.DiscordVerifyRequest) <-chan *api.DiscordVerifyRequest {
	verifiedCh := make(chan *api.DiscordVerifyRequest)

	go func() {
		defer close(verifiedCh)

		res, err := repository.DB.Exec("SELECT ->verified_with FROM ONLY discord_member:$", surgo.ID{input.ID})
		if err != nil {
			s.errCh <- err
			return
		}

		if errors.As(res[0].Error, &surrealdb.ErrNoRow) {
			verifiedCh <- input
		} else {
			s.errCh <- echo.NewHTTPError(http.StatusForbidden, "user is already verified")
		}
	}()

	return verifiedCh
}

func (s *discordService) StrToInt(input string) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		i, err := strconv.Atoi(input)
		if err != nil {
			s.errCh <- echo.NewHTTPError(http.StatusBadRequest, "invalid page number")
			return
		}
		out <- i
	}()

	return out
}

func (s *discordService) GetLeaderboard(page <-chan int) <-chan api.DiscordLeaderboardResponse {
	leaderboardCh := make(chan api.DiscordLeaderboardResponse)

	go func() {
		defer close(leaderboardCh)

		members := make([]model.DiscordMember, 10)
		err := repository.DB.Scan(&members, "SELECT * FROM ONLY discord_member ORDER BY xp DESC LIMIT 10 OFFSET $1", 10*<-page)
		if err != nil {
			s.errCh <- err
			return
		}

		var res api.DiscordLeaderboardResponse
		for _, m := range members {
			res = append(res, api.DiscordLeaderboardEntry{
				DiscordID: m.DiscordID,
				Level:     m.Level,
				XP:        m.XP,
			})
		}

		leaderboardCh <- res
	}()

	return leaderboardCh
}

func (s *discordService) Revoke(id string) <-chan *api.DiscordMemberResponse {
	out := make(chan *api.DiscordMemberResponse)

	go func() {
		defer close(out)

		var member model.DiscordMember
		err := repository.DB.Scan(&member, "DELETE ONLY discord_member:$ RETURN BEFORE", surgo.ID{id})
		if err != nil {
			s.errCh <- err
			return
		}
		out <- &api.DiscordMemberResponse{
			DiscordMember: member,
		}
	}()

	return out
}
