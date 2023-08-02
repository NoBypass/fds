package repository

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"server/db/mappers"
	"server/db/models"
	"server/utils"
)

func FindDiscordByDiscordId(ctx context.Context, driver neo4j.DriverWithContext, discordIdInput string) (*models.Discord, error) {
	result, err := neo4j.ExecuteQuery(ctx, driver,
		"MATCH (d:Discord { discord_id: $discord_id }) RETURN d",
		map[string]any{
			"discord_id": discordIdInput,
		}, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}

	return mappers.ResultToDiscord(result)
}

func CreateDiscord(ctx context.Context, driver neo4j.DriverWithContext, discord *models.DiscordDto) (*models.Discord, error) {
	result, err := neo4j.ExecuteQuery(ctx, driver,
		"CREATE (d:Discord { id: $id, discord_id: $discord_id, name: $name, level: $level, xp: $xp, streak: $streak, last_daily_at: $last_daily_at }) RETURN d",
		map[string]any{
			"id":            utils.GenerateUUID(discord.DiscordID, discord.Name),
			"discord_id":    discord.DiscordID,
			"name":          discord.Name,
			"level":         0,
			"xp":            0,
			"streak":        0,
			"last_daily_at": 0,
		}, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}

	return mappers.ResultToDiscord(result)
}
