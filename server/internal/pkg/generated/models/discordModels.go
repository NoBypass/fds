package models

// Code automatically generated; DO NOT EDIT.

type CreateDiscordInput struct {
	DiscordId string `json:"discordId"`
	Name      string `json:"name"`
}

type GiveXpInput struct {
	DiscordId string `json:"discordId"`
	Amount    int64  `json:"amount"`
}

type DiscordInput struct {
	DiscordId string `json:"discordId"`
}

type Discord struct {
	DiscordId    string        `json:"discord_id"`
	Joined       bool          `json:"joined"`
	VerifiedWith *VerifiedWith `json:"verified_with"`
	LastDailyAt  int64         `json:"last_daily_at"`
	Name         string        `json:"name"`
	Streak       int64         `json:"streak"`
	Level        int64         `json:"level"`
	Xp           int64         `json:"xp"`
}
