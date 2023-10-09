package resolvers

import (
	"context"
	"fmt"
	"server/internal/app/global"
	"server/internal/pkg/auth"
	"server/internal/pkg/generated/models"
	"server/internal/pkg/misc"
	"server/pkg/ogm"
)

func CreateDiscordMutation(ctx context.Context, input *models.CreateDiscordInput) (*models.Discord, error) {
	records, err := global.Get().DB.Query("CREATE (d:Discord { id: $id, discord_id: $discord_id, name: $name, level: $level, xp: $xp, streak: $streak, last_daily_at: $last_daily_at }) RETURN d",
		map[string]any{
			"id":            misc.GenerateUUID(input.DiscordId, input.Name),
			"discord_id":    input.DiscordId,
			"name":          input.Name,
			"level":         0,
			"xp":            0,
			"streak":        0,
			"last_daily_at": 0,
		})
	if err != nil {
		return nil, err
	}
	return ogm.Map(&models.Discord{}, records, "d")
}

func DiscordQuery(ctx context.Context, input *models.DiscordInput) (*models.Discord, error) {
	records, err := global.Get().DB.Query("MATCH (d:Discord { discord_id: $discord_id }) RETURN d", map[string]any{
		"discord_id": input.DiscordId,
	})
	if err != nil {
		return nil, err
	}
	return ogm.Map(&models.Discord{}, records, "d")
}

func GiveXpMutation(ctx context.Context, input *models.GiveXpInput) (*models.Discord, error) {
	err := auth.Allow(ctx, []string{"admin", "bot"})
	if err != nil {
		return nil, err
	}

	db := global.Get().DB
	records, err := db.Query("MATCH (d:Discord { discord_id: $discord_id }) SET d.xp = d.xp + $xp RETURN d", map[string]any{
		"discord_id": input.DiscordId,
		"xp":         input.Amount,
	})
	if err != nil {
		return nil, err
	}

	if records == nil || len(records) == 0 {
		return nil, fmt.Errorf("could not find discord with id %s", input.DiscordId)
	}
	discord, err := ogm.Map(&models.Discord{}, records, "d")
	if err != nil {
		return nil, err
	}

	levelMaxXp := misc.MaxOutAt(discord.Level*1000, 10000)
	if discord.Xp >= levelMaxXp {
		discord.Level += 1
		discord.Xp = discord.Xp - levelMaxXp
	}

	records, err = db.Query("MATCH (d:Discord { discord_id: $discord_id }) SET d.level = $level, d.xp = $xp, d.streak = $streak, d.last_daily_at = $last_daily_at RETURN d",
		map[string]any{
			"discord_id":    discord.DiscordId,
			"level":         discord.Level,
			"xp":            discord.Xp,
			"streak":        discord.Streak,
			"last_daily_at": discord.LastDailyAt,
		})
	if err != nil {
		return nil, err
	}
	discord, err = ogm.Map(&models.Discord{}, records, "d")
	if err != nil {
		return nil, err
	}

	return discord, nil
}
