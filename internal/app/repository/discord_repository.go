package repository

import (
	"github.com/NoBypass/fds/internal/app/errs"
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/surrealdb/surrealdb.go"
	"math"
	"math/rand"
	"time"
)

type DiscordRepository interface {
	Create(input *model.DiscordSignupInput) error
	ClaimDaily(id string) (*model.DiscordMember, error)
}

type discordRepository struct {
	*surrealdb.DB
}

func NewDiscordRepository(db *surrealdb.DB) DiscordRepository {
	return &discordRepository{
		db,
	}
}

func (r *discordRepository) Create(input *model.DiscordSignupInput) error {
	discordMember := model.DiscordMember{
		Nick: input.Nick,
	}
	_, err := r.DB.Create("discord_member:"+input.ID, &discordMember)
	if err != nil {
		return err
	}
	return nil
}

func (r *discordRepository) ClaimDaily(id string) (*model.DiscordMember, error) {
	member, err := surrealdb.SmartUnmarshal[model.DiscordMember](r.DB.Select("discord_member:" + id))
	if err != nil {
		return nil, err
	}

	if member.CanClaimDaily() {
		member.AddXP(math.Round(rand.Float64() * 500.0 * (1.0 + float64(member.Streak)*0.1)))
		member.LastDailyClaim = time.Now().UnixMilli()
		member.Streak++
	} else {
		return nil, errs.NewClaimedError()
	}

	_, err = r.DB.Update("discord_member:"+id, &member)
	if err != nil {
		return nil, err
	}
	return &member, nil
}
