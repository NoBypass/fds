package repository

import (
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/surrealdb/surrealdb.go"
)

type DiscordRepository interface {
	Create(member *model.DiscordMember, profile *model.MojangProfile) error
	Get(id string) (*model.DiscordMember, error)
	Update(member *model.DiscordMember) error
}

type discordRepository struct {
	*surrealdb.DB
}

func NewDiscordRepository(db *surrealdb.DB) DiscordRepository {
	return &discordRepository{
		db,
	}
}

func (r *discordRepository) Create(member *model.DiscordMember, profile *model.MojangProfile) error {
	member.ID = "discord_member:" + member.ID
	_, err := surrealdb.SmartMarshal(r.DB.Create, member)
	if err != nil {
		return err
	}
	_, err = r.DB.Query("RELATE discord_member: $member, mojang_profile: $profile", map[string]string{
		"profile": profile.ID,
		"member":  member.ID,
	})
	return err
}

func (r *discordRepository) Get(id string) (*model.DiscordMember, error) {
	member, err := surrealdb.SmartUnmarshal[model.DiscordMember](r.DB.Select("discord_member:" + id))
	return &member, err
}

func (r *discordRepository) Update(member *model.DiscordMember) error {
	_, err := surrealdb.SmartMarshal(r.DB.Update, member)
	return err
}
