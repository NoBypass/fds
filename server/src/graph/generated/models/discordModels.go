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

//type CreateDiscordInput struct {
//	DiscordId string `json:"discordId"`
//	Name string `json:"name"`
//}
//
//type GiveXpInput struct {
//	DiscordId string `json:"discordId"`
//	Amount int64 `json:"amount"`
//}

type Discord struct {
	DiscordId   string `json:"discord_id"`
	Name        string `json:"name"`
	Level       int64  `json:"level"`
	Xp          int64  `json:"xp"`
	Streak      int64  `json:"streak"`
	LastDailyAt int64  `json:"last_daily_at"`
	//	LinkedWith LinkedWith `json:"linked_with"`
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

	//linkedWith, err := ResultToLinkedWith(result)
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
		// LinkedWith:  *linkedWith,
	}, nil
}

type LinkedWith struct {
	LinkedAt int64   `json:"linked_at"`
	Player   Player  `json:"player"`
	Discord  Discord `json:"discord"`
}

func ResultToLinkedWith(result *neo4j.EagerResult) (*LinkedWith, error) {
	r, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "l")
	if err != nil {
		return nil, err
	}

	linkedAt, err := neo4j.GetProperty[int64](r, "linked_at")
	if err != nil {
		return nil, err
	}

	player, err := ResultToPlayer(result)
	if err != nil {
		return nil, err
	}

	discord, err := ResultToDiscord(result)
	if err != nil {
		return nil, err
	}
	return &LinkedWith{
		LinkedAt: linkedAt,
		Player:   *player,
		Discord:  *discord,
	}, nil
}
