package service

import (
	"fmt"
	"github.com/NoBypass/fds/internal/backend/database"
	"github.com/NoBypass/fds/internal/external"
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/NoBypass/fds/internal/pkg/utils"
	"github.com/NoBypass/mincache"
	"github.com/NoBypass/surgo"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type DiscordService interface {
	Service
	PersistProfile(<-chan model.MojangProfile) <-chan string
	PersistMember(<-chan model.DiscordMember, <-chan struct{})
	PersistPlayer(<-chan model.HypixelPlayer)
	RelateMemberToPlayer(<-chan model.DiscordMember, <-chan model.HypixelPlayer)

	CheckIfAlreadyVerified(*model.DiscordVerifyRequest) <-chan *model.DiscordVerifyRequest
	VerifyHypixelSocials(<-chan model.DiscordMember, <-chan model.HypixelPlayerResponse) (*utils.Broadcaster[model.HypixelPlayer], <-chan struct{})
	FetchHypixelPlayer(*model.DiscordVerifyRequest) (<-chan model.HypixelPlayerResponse, *utils.Broadcaster[model.DiscordMember])
	GiveDaily(<-chan model.DiscordMember) <-chan model.DiscordMember
	GetJWT(*model.DiscordBotLoginRequest) <-chan string
	GetMember(id string) <-chan model.DiscordMember
	StrToInt(string) <-chan int
	GetLeaderboard(page <-chan int) <-chan []model.DiscordLeaderboardEntry
	Revoke(id string) <-chan *model.DiscordMember
}

type discordService struct {
	service
	config        *utils.Config
	hypixelClient *external.HypixelAPIClient
	cache         *mincache.Cache
	db            *surgo.DB
	database.Client
}

func NewDiscordService(config *utils.Config, hypixelClient *external.HypixelAPIClient, db database.Client) DiscordService {
	return &discordService{
		hypixelClient: hypixelClient,
		config:        config,
		Client:        db,
	}
}

func (s *discordService) GetMember(id string) <-chan model.DiscordMember {
	memberCh := make(chan model.DiscordMember)

	s.Pipeline(func(start func() opentracing.Span) error {
		sp := start()

		var member model.DiscordMember
		err := s.DB(sp).Scan(&member, "SELECT * FROM ONLY discord_member:$", surgo.ID{id})
		if err != nil {
			return err
		}

		memberCh <- member
		close(memberCh)
		return nil
	}, s.GetMember)

	return memberCh
}

func (s *discordService) GiveDaily(memberCh <-chan model.DiscordMember) <-chan model.DiscordMember {
	out := make(chan model.DiscordMember)

	s.Pipeline(func(start func() opentracing.Span) error {
		member := <-memberCh
		sp := start()

		lastDaily, err := time.Parse(time.RFC3339, member.LastDailyAt)
		if err != nil {
			return err
		} else if lastDaily.After(time.Now().Truncate(time.Hour * 24)) {
			return echo.NewHTTPError(http.StatusConflict, "already claimed daily today")
		}

		member.AddXP(math.Round(rand.Float64() * 500.0 * (1.0 + float64(member.Streak)*0.1)))
		member.LastDailyAt = time.Now().Format(time.RFC3339)
		member.Streak++

		var newMember model.DiscordMember
		err = s.DB(sp).Scan(&newMember, "UPDATE ONLY discord_member:$ MERGE {"+
			"last_daily_at: $last_daily_at,"+
			"xp: $xp,"+
			"streak: $streak,"+
			"level: $level"+
			"} RETURN AFTER", surgo.ID{member.DiscordID}, member)
		if err != nil {
			return err
		}

		out <- newMember
		close(out)
		return nil
	}, s.GiveDaily)

	return out
}

func (s *discordService) GetJWT(input *model.DiscordBotLoginRequest) <-chan string {
	tokenCh := make(chan string)

	s.Pipeline(func(start func() opentracing.Span) error {
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

		close(tokenCh)
		return nil
	}, s.GetJWT)

	return tokenCh
}

func (s *discordService) FetchHypixelPlayer(input *model.DiscordVerifyRequest) (<-chan model.HypixelPlayerResponse, *utils.Broadcaster[model.DiscordMember]) {
	playerCh := make(chan model.HypixelPlayerResponse)
	memberCh := make(chan model.DiscordMember)
	memberBr := utils.NewBroadcaster(memberCh)

	s.Pipeline(func(start func() opentracing.Span) error {
		sp := start()
		player, err := s.hypixelClient.Player(input.Nick, sp)
		if err != nil {
			return err
		}

		if !player.Success {
			ext.LogError(sp, fmt.Errorf("%+v", player))
			return echo.NewHTTPError(http.StatusNotFound, "hypixel: player not found")
		}

		playerCh <- *player
		memberCh <- model.DiscordMember{
			DiscordID:   input.ID,
			Name:        input.Name,
			Nick:        player.Player.DisplayName,
			LastDailyAt: time.Now().Add(-time.Hour * 24).Format(time.RFC3339),
		}

		close(playerCh)
		return nil
	}, s.FetchHypixelPlayer)

	return playerCh, memberBr
}

func (s *discordService) VerifyHypixelSocials(memberCh <-chan model.DiscordMember, playerCh <-chan model.HypixelPlayerResponse) (*utils.Broadcaster[model.HypixelPlayer], <-chan struct{}) {
	outPlayerCh := make(chan model.HypixelPlayer)
	awaitVerify := make(chan struct{})
	playerBr := utils.NewBroadcaster(outPlayerCh)

	s.Pipeline(func(start func() opentracing.Span) error {
		player := <-playerCh
		member := <-memberCh
		sp := start()

		if player.Player.SocialMedia.Links.Discord == member.Name {
			var exists bool
			err := s.DB(sp).Scan(&exists, `
			RETURN (
				SELECT * FROM (
					SELECT <-verified_with<-discord_member AS member FROM (
						SELECT * FROM hypixel_player 
						WHERE string::lowercase($1)=string::lowercase(name) 
						ORDER BY date DESC 
						LIMIT 1
					) FETCH member
				).member[0]
			)[0].discord_id=$2;`, player.Player.DisplayName, member.DiscordID)
			if err != nil {
				return err
			}

			if !exists {
				awaitVerify <- struct{}{}
				outPlayerCh <- model.HypixelPlayer{
					UUID: player.Player.UUID,
					Date: time.Now().Format(time.RFC3339),
					Name: player.Player.DisplayName,
				}
			} else {
				return echo.NewHTTPError(http.StatusConflict, "already verified")
			}
		} else {
			return echo.NewHTTPError(http.StatusForbidden, "discord tag does not match hypixel socials")
		}

		close(outPlayerCh)
		close(awaitVerify)
		return nil
	}, s.VerifyHypixelSocials)

	return playerBr, awaitVerify
}

func (s *discordService) PersistProfile(profileCh <-chan model.MojangProfile) <-chan string {
	actual := make(chan string)

	s.Pipeline(func(start func() opentracing.Span) error {
		profile := <-profileCh
		sp := start()

		res, err := s.DB(sp).Exec("CREATE mojang_profile:$ CONTENT {"+
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
		close(actual)
		return nil
	}, s.PersistProfile)

	return actual
}

func (s *discordService) PersistMember(memberCh <-chan model.DiscordMember, awaitVerify <-chan struct{}) {
	s.Pipeline(func(start func() opentracing.Span) error {
		member := <-memberCh
		<-awaitVerify
		sp := start()

		res, err := s.DB(sp).Exec("CREATE discord_member:$ CONTENT {"+
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

		res, err := s.DB(sp).Exec("CREATE hypixel_player:$ CONTENT {"+
			"uuid: $uuid,"+
			"date: $date,"+
			"name: $name"+
			"}", player, surgo.ID{player.Name, player.Date})
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

		res, err := s.DB(sp).Exec("RELATE discord_member:$->verified_with->hypixel_player:$", surgo.ID{member.DiscordID}, surgo.ID{player.Name, player.Date})
		if err != nil {
			return err
		}
		if err := utils.Error(res); err != nil {
			return err
		}

		return nil
	}, s.RelateMemberToPlayer)
}

func (s *discordService) CheckIfAlreadyVerified(input *model.DiscordVerifyRequest) <-chan *model.DiscordVerifyRequest {
	verifiedCh := make(chan *model.DiscordVerifyRequest)

	s.Pipeline(func(start func() opentracing.Span) error {
		sp := start()

		res, err := s.DB(sp).Exec("SELECT ->verified_with FROM discord_member:$", surgo.ID{input.ID})
		if err != nil {
			return err
		}
		if len(res[0].Data.([]any)) == 0 {
			verifiedCh <- input
		} else {
			return echo.NewHTTPError(http.StatusConflict, "already verified")
		}

		close(verifiedCh)
		return nil
	}, s.CheckIfAlreadyVerified)

	return verifiedCh
}

func (s *discordService) StrToInt(input string) <-chan int {
	out := make(chan int)

	s.Pipeline(func(start func() opentracing.Span) error {
		start()

		i, err := strconv.Atoi(input)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid page number")
		}
		out <- i

		close(out)
		return nil
	}, s.StrToInt)

	return out
}

func (s *discordService) GetLeaderboard(page <-chan int) <-chan []model.DiscordLeaderboardEntry {
	leaderboardCh := make(chan []model.DiscordLeaderboardEntry)

	s.Pipeline(func(start func() opentracing.Span) error {
		sp := start()

		members := make([]model.DiscordLeaderboardEntry, 10)
		err := s.DB(sp).Scan(&members, "SELECT * FROM discord_member ORDER BY xp DESC LIMIT 10 START $1", 10*<-page)
		if err != nil {
			return err
		}

		leaderboardCh <- members
		close(leaderboardCh)
		return nil
	}, s.GetLeaderboard)

	return leaderboardCh
}

func (s *discordService) Revoke(id string) <-chan *model.DiscordMember {
	out := make(chan *model.DiscordMember)

	s.Pipeline(func(start func() opentracing.Span) error {
		sp := start()

		var member model.DiscordMember
		_, err := s.DB(sp).Exec("DELETE person:tobie->verified_with", surgo.ID{id})
		if err != nil {
			return err
		}
		err = s.DB(sp).Scan(&member, "DELETE ONLY discord_member:$ RETURN BEFORE", surgo.ID{id})
		if err != nil {
			return err
		}
		out <- &member

		close(out)
		return nil
	}, s.Revoke)

	return out
}
