package mappers

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"server/db/models"
)

func ResultToDiscord(result *neo4j.EagerResult) (*models.Discord, error) {
	discordNode, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "d")
	if err != nil {
		return nil, err
	}

	id, err := neo4j.GetProperty[string](discordNode, "id")
	if err != nil {
		return nil, err
	}

	discordId, err := neo4j.GetProperty[string](discordNode, "discord_id")
	if err != nil {
		return nil, err
	}

	name, err := neo4j.GetProperty[string](discordNode, "name")
	if err != nil {
		return nil, err
	}

	level, err := neo4j.GetProperty[int64](discordNode, "level")
	if err != nil {
		return nil, err
	}

	xp, err := neo4j.GetProperty[int64](discordNode, "xp")
	if err != nil {
		return nil, err
	}

	streak, err := neo4j.GetProperty[int64](discordNode, "streak")
	if err != nil {
		return nil, err
	}

	lastDailyAt, err := neo4j.GetProperty[int64](discordNode, "last_daily_at")
	if err != nil {
		return nil, err
	}

	return &models.Discord{
		ID:          id,
		DiscordID:   discordId,
		Name:        name,
		Level:       level,
		Xp:          xp,
		Streak:      streak,
		LastDailyAt: lastDailyAt,
	}, nil
}
