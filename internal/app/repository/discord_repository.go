package repository

import (
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/NoBypass/fds/internal/pkg/surreal_wrap"
	"github.com/surrealdb/surrealdb.go"
	"time"
)

type DiscordRepository interface {
	Repository
	Create(member *model.DiscordMember) error
	Get(id string) (*model.DiscordMember, error)
	Update(id string, member *model.DiscordMember) error

	RelatePlayedWith(in *model.MojangProfile, out *model.HypixelPlayer) error
	RelateVerifiedWith(in *model.DiscordMember, out *model.HypixelPlayer) error
}

type discordRepository struct {
	repository
}

func NewDiscordRepository(db *surreal_wrap.DB) DiscordRepository {
	return &discordRepository{
		newRepository(db),
	}
}

func (r *discordRepository) Create(member *model.DiscordMember) error {
	if member.LastDailyAt == "" {
		member.LastDailyAt = time.Now().Add(-24 * time.Hour).Format(time.RFC3339)
	}
	_, err := r.DB.Queryf(`CREATE discord_member:%s CONTENT {
		"discord_id": "%s",
		"name": "%s",
		"nick": "%s",
		"xp": %f,
		"level": %d,
		"streak": %d,
		"last_daily_at": "%s"
	}`, member.DiscordID, member.DiscordID, member.Name, member.Nick, member.XP, member.Level, member.Streak, member.LastDailyAt)
	return err
}

func (r *discordRepository) Get(id string) (*model.DiscordMember, error) {
	member, err := surrealdb.SmartUnmarshal[model.DiscordMember](r.DB.Select("discord_member:" + id))
	return &member, err
}

func (r *discordRepository) Update(id string, member *model.DiscordMember) error {
	_, err := r.DB.Queryf(`UPDATE discord_member:%s SET {
		"name": "%s",
		"nick": "%s",
		"xp": %f,
		"level": %d,
		"streak": %d,
		"last_daily_at": "%s"
	}`, id, member.Name, member.Nick, member.XP, member.Level, member.Streak, member.LastDailyAt)
	return err
}

func (r *discordRepository) RelatePlayedWith(in *model.MojangProfile, out *model.HypixelPlayer) error {
	_, err := r.DB.Queryf(`RELATE mojang_profile:["%s", "%s"]->played_with->hypixel_player:["%s", "%s"]`, in.UUID, in.Date, out.UUID, out.Date)
	return err
}

func (r *discordRepository) RelateVerifiedWith(in *model.DiscordMember, out *model.HypixelPlayer) error {
	_, err := r.DB.Queryf(`RELATE discord_member:%s->verified_with->hypixel_player:["%s", "%s"]`, in.DiscordID, out.UUID, out.Date)
	return err
}
