package repository

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"server/src/graph/generated"
	"server/src/utils"
)

func FindDiscordByDiscordId(ctx context.Context, driver neo4j.DriverWithContext, discordIdInput string) (*neo4j.EagerResult, error) {
	result, err := neo4j.ExecuteQuery(ctx, driver,
		"MATCH (d:Discord { discord_id: $discord_id }) RETURN d",
		map[string]any{
			"discord_id": discordIdInput,
		}, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func CreateDiscord(ctx context.Context, driver neo4j.DriverWithContext, discord *generated.Discord) (*neo4j.EagerResult, error) {
	result, err := neo4j.ExecuteQuery(ctx, driver,
		"CREATE (d:Discord { id: $id, discord_id: $discord_id, name: $name, level: $level, xp: $xp, streak: $streak, last_daily_at: $last_daily_at }) RETURN d",
		map[string]any{
			"id":            utils.GenerateUUID(discord.DiscordId, discord.Name),
			"discord_id":    discord.DiscordId,
			"name":          discord.Name,
			"level":         0,
			"xp":            0,
			"streak":        0,
			"last_daily_at": 0,
		}, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}

	return result, nil
}