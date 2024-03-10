package api

import (
	"github.com/NoBypass/fds/internal/pkg/model"
)

type DiscordVerifyRequest struct {
	ID   string `json:"id"`
	Nick string `json:"nick"`
	Name string `json:"name"`
}

type DiscordVerifyResponse struct {
	Actual string `json:"actual"`
}

type DiscordMemberResponse struct {
	model.DiscordMember
}

type DiscordBotLoginRequest struct {
	Pwd string `json:"pwd" query:"pwd"`
}

type DiscordBotLoginResponse struct {
	Token string `json:"token"`
}

type DiscordLeaderboardResponse []DiscordLeaderboardEntry

type DiscordLeaderboardEntry struct {
	DiscordID string  `json:"discord_id"`
	Level     int     `json:"level"`
	XP        float64 `json:"xp"`
}
