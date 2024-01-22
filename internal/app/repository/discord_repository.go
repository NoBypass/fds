package repository

import (
	"github.com/NoBypass/fds/internal/app/custom_err"
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/surrealdb/surrealdb.go"
	"math"
	"math/rand"
	"time"
)

type DiscordRepository interface {
	Create(input *model.MojangResponse) error
	ClaimDaily(id string) (*model.DiscordDailyResponse, error)
}

type discordRepository struct {
	*surrealdb.DB
}

func NewDiscordRepository(db *surrealdb.DB) DiscordRepository {
	return &discordRepository{
		db,
	}
}

func (r *discordRepository) Create(input *model.MojangResponse) error {
	discordMember := model.DiscordMember{
		Nick: input.Name,
	}
	_, err := r.DB.Create("discord_member:"+input.ID, &discordMember)
	if err != nil {
		return err
	}
	return nil
}

func (r *discordRepository) ClaimDaily(id string) (*model.DiscordDailyResponse, error) {
	member, err := surrealdb.SmartUnmarshal[model.DiscordMember](r.DB.Select("discord_member:" + id))
	if err != nil {
		return nil, err
	}

	oldLvl := member.Level
	gain := math.Round(rand.Float64() * 500.0)
	withBonus := gain * (1.0 + float64(member.Streak)*0.1)
	if member.CanClaimDaily() {
		member.AddXP(withBonus)
		member.LastDailyClaim = time.Now().UnixMilli()
		member.Streak++
	} else {
		return nil, custom_err.NewClaimedError()
	}

	_, err = r.DB.Update("discord_member:"+id, &member)
	if err != nil {
		return nil, err
	}
	return &model.DiscordDailyResponse{
		XP:        member.XP,
		Level:     member.Level,
		Levelup:   oldLvl != member.Level,
		Needed:    member.GetNeededXP(),
		Streak:    member.Streak,
		WithBonus: withBonus,
		Gained:    gain,
	}, nil
}
