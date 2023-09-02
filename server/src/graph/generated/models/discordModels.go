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
	DiscordId   string `json:"discord_id"`
	Name        string `json:"name"`
	Level       int64  `json:"level"`
	Xp          int64  `json:"xp"`
	Streak      int64  `json:"streak"`
	LastDailyAt int64  `json:"last_daily_at"`
}

func ResultToDiscord(result *neo4j.EagerResult) (*Discord, error) {
	r, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "d")
	if err != nil {
		return nil, err
	}

	discordId, err := neo4j.GetProperty[string](r, "discord_id")
	if err != nil {
		return nil, err
	}

	name, err := neo4j.GetProperty[string](r, "name")
	if err != nil {
		return nil, err
	}

	level, err := neo4j.GetProperty[int64](r, "level")
	if err != nil {
		return nil, err
	}

	xp, err := neo4j.GetProperty[int64](r, "xp")
	if err != nil {
		return nil, err
	}

	streak, err := neo4j.GetProperty[int64](r, "streak")
	if err != nil {
		return nil, err
	}

	lastDailyAt, err := neo4j.GetProperty[int64](r, "last_daily_at")
	if err != nil {
		return nil, err
	}
	return &Discord{
		DiscordId:   discordId,
		Name:        name,
		Level:       level,
		Xp:          xp,
		Streak:      streak,
		LastDailyAt: lastDailyAt,
	}, nil
}
