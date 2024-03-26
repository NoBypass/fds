package service

import (
	"encoding/json"
	"fmt"
	"github.com/NoBypass/fds/internal/backend/repository"
	"github.com/NoBypass/fds/internal/pkg/conf"
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/NoBypass/fds/internal/pkg/utils"
	"github.com/NoBypass/fds/pkg/api"
	"github.com/NoBypass/surgo"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type DiscordService interface {
	Service
	PersistProfile(<-chan model.MojangProfile) <-chan string
	PersistMember(<-chan model.DiscordMember)
	PersistPlayer(<-chan model.HypixelPlayer)
	RelateMemberToPlayer(<-chan model.DiscordMember, <-chan model.HypixelPlayer)

	CheckIfAlreadyVerified(*api.DiscordVerifyRequest) <-chan *api.DiscordVerifyRequest
	VerifyHypixelSocials(<-chan model.DiscordMember, <-chan model.HypixelPlayerResponse) *utils.Broadcaster[model.HypixelPlayer]
	FetchHypixelPlayer(<-chan *api.DiscordVerifyRequest) (*utils.Broadcaster[model.HypixelPlayerResponse], *utils.Broadcaster[model.DiscordMember])
	GiveDaily(<-chan model.DiscordMember) <-chan model.DiscordMember
	GetJWT(*api.DiscordBotLoginRequest) <-chan string
	GetMember(id string) <-chan model.DiscordMember
	StrToInt(string) <-chan int
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

	s.Pipeline(func(start func() opentracing.Span) error {
		defer close(memberCh)
		sp := start()

		member := new(model.DiscordMember)
		err := repository.DB(sp).Scan(member, "SELECT * FROM ONLY discord_member:$", surgo.ID{id})
		if err != nil {
			return err
		}

		memberCh <- *member

		return nil
	}, s.GetMember)

	return memberCh
}

func (s *discordService) GiveDaily(memberCh <-chan model.DiscordMember) <-chan model.DiscordMember {
	out := make(chan model.DiscordMember)

	s.Pipeline(func(start func() opentracing.Span) error {
		defer close(out)

		member := <-memberCh
		sp := start()

		member.AddXP(math.Round(rand.Float64() * 500.0 * (1.0 + float64(member.Streak)*0.1)))
		member.LastDailyAt = time.Now().Add(-time.Hour * 24).Format(time.RFC3339)
		member.Streak++

		var newMember model.DiscordMember
		err := repository.DB(sp).Scan(&newMember, "UPDATE discord_member:$ MERGE {"+
			"last_daily_at: $last_daily_at,"+
			"xp: $xp,"+
			"streak: $streak"+
			"}", member, surgo.ID{member.DiscordID})
		if err != nil {
			return err
		}
		out <- newMember

		return nil
	}, s.GiveDaily)

	return memberCh
}

func (s *discordService) GetJWT(input *api.DiscordBotLoginRequest) <-chan string {
	tokenCh := make(chan string)

	s.Pipeline(func(start func() opentracing.Span) error {
		defer close(tokenCh)
		start()

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
				return err
			}
			tokenCh <- signedToken
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid password")
		}

		return nil
	}, s.GetJWT)

	return tokenCh
}

func (s *discordService) FetchHypixelPlayer(inputCh <-chan *api.DiscordVerifyRequest) (*utils.Broadcaster[model.HypixelPlayerResponse], *utils.Broadcaster[model.DiscordMember]) {
	playerCh := make(chan model.HypixelPlayerResponse)
	memberCh := make(chan model.DiscordMember)
	playerBr := utils.NewBroadcaster(playerCh)
	memberBr := utils.NewBroadcaster(memberCh)

	s.Pipeline(func(start func() opentracing.Span) error {
		defer close(playerCh)

		input := <-inputCh
		start()
		req, err := http.NewRequest(http.MethodGet, "https://api.hypixel.net/player?name="+input.Nick, nil)
		if err != nil {
			return err
		}

		req.Header.Add("API-Key", s.config.HypixelAPIKey)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return fmt.Errorf("hypixel: %s", err)
		}

		var player model.HypixelPlayerResponse
		err = json.NewDecoder(resp.Body).Decode(&player)
		if err != nil {
			return err
		}

		playerCh <- player
		memberCh <- model.DiscordMember{
			DiscordID:   input.ID,
			Name:        input.Name,
			Nick:        player.Player.DisplayName,
			LastDailyAt: time.Now().Add(-time.Hour * 24).Format(time.RFC3339),
		}
		return nil
	}, s.FetchHypixelPlayer)

	return playerBr, memberBr
}

func (s *discordService) VerifyHypixelSocials(memberCh <-chan model.DiscordMember, playerCh <-chan model.HypixelPlayerResponse) *utils.Broadcaster[model.HypixelPlayer] {
	outPlayerCh := make(chan model.HypixelPlayer)
	playerBr := utils.NewBroadcaster(outPlayerCh)

	s.Pipeline(func(start func() opentracing.Span) error {
		defer close(outPlayerCh)

		member := <-memberCh
		player := <-playerCh
		start()

		if player.Success {
			if player.Player.SocialMedia.Links.Discord == member.Name {
				outPlayerCh <- model.HypixelPlayer{
					UUID: player.Player.UUID,
					Date: time.Now().Format(time.RFC3339),
					Name: player.Player.DisplayName,
				}
			} else {
				return echo.NewHTTPError(http.StatusForbidden, "discord id does not match hypixel socials")
			}
		} else {
			return echo.NewHTTPError(http.StatusNotFound, "hypixel: player not found")
		}

		return nil
	}, s.VerifyHypixelSocials)

	return playerBr
}

func (s *discordService) PersistProfile(profileCh <-chan model.MojangProfile) <-chan string {
	actual := make(chan string)

	s.Pipeline(func(start func() opentracing.Span) error {
		profile := <-profileCh
		sp := start()

		res, err := repository.DB(sp).Exec("CREATE mojang_profile:$ CONTENT {"+
			"date: $date,"+
			"uuid: $uuid,"+
			"name: $name"+
			"}", profile, surgo.ID{profile.UUID, profile.Date})
		if err != nil {
			return err
		}
		if err := utils.Error(res); err != nil {
			return err
		}

		actual <- profile.Name
		return nil
	}, s.PersistProfile)

	return actual
}

func (s *discordService) PersistMember(memberCh <-chan model.DiscordMember) {
	s.Pipeline(func(start func() opentracing.Span) error {
		member := <-memberCh
		sp := start()

		res, err := repository.DB(sp).Exec("CREATE discord_member:$ CONTENT {"+
			"discord_id: $discord_id,"+
			"name: $name,"+
			"nick: $nick,"+
			"xp: 0,"+
			"last_daily_at: $last_daily_at,"+
			"level: 0,"+
			"streak: 0"+
			"}", member, surgo.ID{member.DiscordID})
		if err != nil {
			return err
		}
		if err := utils.Error(res); err != nil {
			return err
		}

		return nil
	}, s.PersistMember)
}

func (s *discordService) PersistPlayer(playerCh <-chan model.HypixelPlayer) {
	s.Pipeline(func(start func() opentracing.Span) error {
		player := <-playerCh
		sp := start()

		res, err := repository.DB(sp).Exec("CREATE hypixel_player:$ CONTENT {"+
			"uuid: $uuid,"+
			"date: $date,"+
			"name: $name"+
			"}", player, surgo.ID{player.UUID, player.Date})
		if err != nil {
			return err
		}
		if err := utils.Error(res); err != nil {
			return err
		}

		return nil
	}, s.PersistPlayer)
}

func (s *discordService) RelateMemberToPlayer(memberCh <-chan model.DiscordMember, playerCh <-chan model.HypixelPlayer) {
	s.Pipeline(func(start func() opentracing.Span) error {
		member := <-memberCh
		player := <-playerCh
		sp := start()

		res, err := repository.DB(sp).Exec("RELATE discord_member:$->verified_with->hypixel_player:$", surgo.ID{member.DiscordID}, surgo.ID{player.UUID})
		if err != nil {
			return err
		}
		if err := utils.Error(res); err != nil {
			return err
		}

		return nil
	}, s.RelateMemberToPlayer)
}

func (s *discordService) CheckIfAlreadyVerified(input *api.DiscordVerifyRequest) <-chan *api.DiscordVerifyRequest {
	verifiedCh := make(chan *api.DiscordVerifyRequest)

	s.Pipeline(func(start func() opentracing.Span) error {
		defer close(verifiedCh)
		sp := start()

		res, err := repository.DB(sp).Exec("SELECT ->verified_with FROM discord_member:$", surgo.ID{input.ID})
		if err != nil {
			return err
		}
		if len(res[0].Data.([]any)) == 0 {
			verifiedCh <- input
		} else {
			return echo.NewHTTPError(http.StatusConflict, "already verified")
		}

		return nil
	}, s.CheckIfAlreadyVerified)

	return verifiedCh
}

func (s *discordService) StrToInt(input string) <-chan int {
	out := make(chan int)

	s.Pipeline(func(start func() opentracing.Span) error {
		defer close(out)
		start()

		i, err := strconv.Atoi(input)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid page number")
		}
		out <- i

		return nil
	}, s.StrToInt)

	return out
}

func (s *discordService) GetLeaderboard(page <-chan int) <-chan api.DiscordLeaderboardResponse {
	leaderboardCh := make(chan api.DiscordLeaderboardResponse)

	s.Pipeline(func(start func() opentracing.Span) error {
		defer close(leaderboardCh)
		sp := start()

		members := make([]model.DiscordMember, 10)
		err := repository.DB(sp).Scan(&members, "SELECT * FROM ONLY discord_member ORDER BY xp DESC LIMIT 10 OFFSET $1", 10*<-page)
		if err != nil {
			return err
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
		return nil
	}, s.GetLeaderboard)

	return leaderboardCh
}

func (s *discordService) Revoke(id string) <-chan *api.DiscordMemberResponse {
	out := make(chan *api.DiscordMemberResponse)

	s.Pipeline(func(start func() opentracing.Span) error {
		defer close(out)
		sp := start()

		var member model.DiscordMember
		err := repository.DB(sp).Scan(&member, "DELETE ONLY discord_member:$ RETURN BEFORE", surgo.ID{id})
		if err != nil {
			return err
		}
		out <- &api.DiscordMemberResponse{
			DiscordMember: member,
		}

		return nil
	}, s.Revoke)

	return out
}
