package services

import "server/internal/pkg/model"

type DiscordService interface {
	Daily(member *model.DiscordMember) error
}
