package ogm

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"server/internal/pkg/generated/models"
)

func GiveXp(ctx context.Context, driver neo4j.DriverWithContext, discordIdInput string, xp int64) (*neo4j.EagerResult, error) {
	result, err := neo4j.ExecuteQuery(ctx, driver,
		"MATCH (d:Discord { discord_id: $discord_id }) SET d.xp = d.xp + $xp RETURN d",
		map[string]any{
			"discord_id": discordIdInput,
			"xp":         xp,
		}, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateDiscord(ctx context.Context, driver neo4j.DriverWithContext, discord *models.Discord) (*neo4j.EagerResult, error) {
	result, err := neo4j.ExecuteQuery(ctx, driver,
		"MATCH (d:Discord { discord_id: $discord_id }) SET d.level = $level, d.xp = $xp, d.streak = $streak, d.last_daily_at = $last_daily_at RETURN d",
		map[string]any{
			"discord_id":    discord.DiscordId,
			"level":         discord.Level,
			"xp":            discord.Xp,
			"streak":        discord.Streak,
			"last_daily_at": discord.LastDailyAt,
		}, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}

	return result, nil
}
