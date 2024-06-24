package service

import (
	"context"
	"github.com/NoBypass/fds/internal/backend/tracing"
	"github.com/NoBypass/fds/internal/external"
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/NoBypass/fds/internal/pkg/utils"
	"github.com/NoBypass/mincache"
	"github.com/NoBypass/surgo"
	"github.com/labstack/echo/v4"
	"math"
	"math/rand"
	"net/http"
	"time"
)

type DiscordService interface {
	GetMember(ctx context.Context, id string) (*model.DiscordMember, error)
	GiveDaily(context.Context, *model.DiscordMember) error
	FetchHypixelPlayer(context.Context, *model.DiscordVerifyRequest) (*model.HypixelPlayerResponse, *model.DiscordMember, error)
	VerifyHypixelSocials(context.Context, *model.DiscordMember, *model.HypixelPlayerResponse) (*model.HypixelPlayer, error)
	Revoke(ctx context.Context, id string) (*model.DiscordMember, error)
	GetLeaderboard(ctx context.Context, page int) ([]model.DiscordLeaderboardEntry, error)

	PersistMember(context.Context, *model.DiscordMember) error
	PersistPlayer(context.Context, *model.HypixelPlayer) error
	RelateMemberToPlayer(context.Context, *model.DiscordMember, *model.HypixelPlayer) error
}

type discordService struct {
	DatabaseService
	tracing.Tracable

	config        *utils.Config
	hypixelClient *external.HypixelAPIClient
	cache         *mincache.Cache
}

func NewDiscordService(config *utils.Config, hypixelClient *external.HypixelAPIClient, db DatabaseService) DiscordService {
	return &discordService{
		hypixelClient:   hypixelClient,
		config:          config,
		DatabaseService: db,
		Tracable:        tracing.NewTracable(),
	}
}

func (s *discordService) GetMember(ctx context.Context, id string) (*model.DiscordMember, error) {
	sp, ctx := s.StartSpan(ctx, s.GetMember)
	defer sp.Finish()

	var member model.DiscordMember
	err := s.DB(sp).Scan(&member, "SELECT * FROM ONLY discord_member:$", surgo.ID{id})
	return &member, err
}

func (s *discordService) GiveDaily(ctx context.Context, member *model.DiscordMember) error {
	sp, ctx := s.StartSpan(ctx, s.GiveDaily)
	defer sp.Finish()

	lastDaily, err := time.Parse(time.RFC3339, member.LastDailyAt)
	if err != nil {
		return err
	} else if lastDaily.After(time.Now().Truncate(time.Hour * 24)) {
		return echo.NewHTTPError(http.StatusConflict, "already claimed daily today")
	}

	member.AddXP(math.Round(rand.Float64() * 500.0 * (1.0 + float64(member.Streak)*0.1)))
	member.LastDailyAt = time.Now().Format(time.RFC3339)
	member.Streak++

	err = s.DB(sp).Scan(&member, "UPDATE ONLY discord_member:$ MERGE {"+
		"last_daily_at: $last_daily_at,"+
		"xp: $xp,"+
		"streak: $streak,"+
		"level: $level"+
		"} RETURN AFTER", surgo.ID{member.DiscordID}, member)
	if err != nil {
		return err
	}

	return nil
}

func (s *discordService) FetchHypixelPlayer(ctx context.Context, req *model.DiscordVerifyRequest) (*model.HypixelPlayerResponse, *model.DiscordMember, error) {
	sp, ctx := s.StartSpan(ctx, s.FetchHypixelPlayer)
	defer sp.Finish()

	player, err := s.hypixelClient.Player(req.Nick, sp)
	if err != nil {
		return nil, nil, err
	}

	if !player.Success {
		return nil, nil, echo.NewHTTPError(http.StatusNotFound, "hypixel: player not found")
	}

	return player, &model.DiscordMember{
		DiscordID:   req.ID,
		Name:        req.Name,
		Nick:        player.Player.DisplayName,
		LastDailyAt: time.Now().Add(-time.Hour * 24).Format(time.RFC3339),
	}, nil
}

func (s *discordService) VerifyHypixelSocials(ctx context.Context, member *model.DiscordMember, resp *model.HypixelPlayerResponse) (*model.HypixelPlayer, error) {
	sp, ctx := s.StartSpan(ctx, s.VerifyHypixelSocials)
	defer sp.Finish()

	if resp.Player.SocialMedia.Links.Discord == member.Name {
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
			)[0].discord_id=$2;`, resp.Player.DisplayName, member.DiscordID)
		if err != nil {
			return nil, err
		} else if exists {
			return nil, echo.NewHTTPError(http.StatusConflict, "already verified")
		}
	} else {
		return nil, echo.NewHTTPError(http.StatusForbidden, "discord tag does not match hypixel socials")
	}

	return &model.HypixelPlayer{
		UUID: resp.Player.UUID,
		Date: time.Now().Format(time.RFC3339),
		Name: resp.Player.DisplayName,
	}, nil
}

func (s *discordService) PersistMember(ctx context.Context, member *model.DiscordMember) error {
	sp, ctx := s.StartSpan(ctx, s.PersistMember)
	defer sp.Finish()

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
	} else if err = utils.Error(res); err != nil {
		return err
	}

	return nil
}

func (s *discordService) PersistPlayer(ctx context.Context, player *model.HypixelPlayer) error {
	sp, ctx := s.StartSpan(ctx, s.PersistPlayer)
	defer sp.Finish()

	res, err := s.DB(sp).Exec("CREATE hypixel_player:$ CONTENT {"+
		"uuid: $uuid,"+
		"date: $date,"+
		"name: $name"+
		"}", player, surgo.ID{player.Name, player.Date})
	if err != nil {
		return err
	} else if err := utils.Error(res); err != nil {
		return err
	}

	return nil
}

func (s *discordService) RelateMemberToPlayer(ctx context.Context, member *model.DiscordMember, player *model.HypixelPlayer) error {
	sp, ctx := s.StartSpan(ctx, s.RelateMemberToPlayer)
	defer sp.Finish()

	res, err := s.DB(sp).Exec("RELATE discord_member:$->verified_with->hypixel_player:$", surgo.ID{member.DiscordID}, surgo.ID{player.Name, player.Date})
	if err != nil {
		return err
	} else if err := utils.Error(res); err != nil {
		return err
	}

	return nil
}

func (s *discordService) GetLeaderboard(ctx context.Context, page int) ([]model.DiscordLeaderboardEntry, error) {
	sp, ctx := s.StartSpan(ctx, s.GetLeaderboard)
	defer sp.Finish()

	members := make([]model.DiscordLeaderboardEntry, 10)
	err := s.DB(sp).Scan(&members, "SELECT * FROM discord_member ORDER BY xp DESC LIMIT 10 START $1", 10*page)
	if err != nil {
		return nil, err
	}

	return members, nil
}

func (s *discordService) Revoke(ctx context.Context, id string) (*model.DiscordMember, error) {
	sp, ctx := s.StartSpan(ctx, s.VerifyHypixelSocials)
	defer sp.Finish()

	var member model.DiscordMember
	_, err := s.DB(sp).Exec("DELETE person:tobie->verified_with", surgo.ID{id})
	if err != nil {
		return nil, err
	}
	err = s.DB(sp).Scan(&member, "DELETE ONLY discord_member:$ RETURN BEFORE", surgo.ID{id})
	if err != nil {
		return nil, err
	}

	return &member, nil
}
