package models

import "github.com/neo4j/neo4j-go-driver/v5/neo4j"

// Code automatically generated; DO NOT EDIT.

type DiscordInput struct {
	DiscordId string `json:"discordId"`
}

type CreateDiscordInput struct {
	DiscordId string `json:"discordId"`
	Name      string `json:"name"`
}

type GiveXpInput struct {
	DiscordId string `json:"discordId"`
	Amount    int64  `json:"amount"`
}

type Discord struct {
	Name         string        `json:"name"`
	LastDailyAt  int64         `json:"last_daily_at"`
	DiscordId    string        `json:"discord_id"`
	Streak       int64         `json:"streak"`
	Level        int64         `json:"level"`
	Xp           int64         `json:"xp"`
	VerifiedWith *VerifiedWith `json:"verified_with"`
}
