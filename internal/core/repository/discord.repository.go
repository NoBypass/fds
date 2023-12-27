package repository

import "server/internal/pkg/model"

func (r *Repository) CreateDiscord(input *model.DiscordMemberInput) error {
	discordMember := model.DiscordMember{
		Nick: input.Nick,
	}
	_, err := r.DB.Create("discord_member:"+input.ID, &discordMember)
	if err != nil {
		return err
	}
	return nil
}
