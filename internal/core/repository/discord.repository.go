package repository

import "server/internal/pkg/model"

type DiscordRepository interface {
	CreateDiscord(input *model.DiscordSignupInput) error
}

func (r *Repository) CreateDiscord(input *model.DiscordSignupInput) error {
	discordMember := model.DiscordMember{
		Nick: input.Nick,
	}
	_, err := r.DB.Create("discord_member:"+input.ID, &discordMember)
	if err != nil {
		return err
	}
	return nil
}
