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
	discords := ogm.New[models.Discord](ctx)

	return discords.Create(&models.Discord{
		DiscordId: input.DiscordId,
		Name:      input.Name,
	})
}

func DiscordQuery(ctx context.Context, input *models.DiscordInput) (*models.Discord, error) {
	discords := ogm.New[models.Discord](ctx)
	return discords.Find(&models.Discord{DiscordId: input.DiscordId})
}

func GiveXpMutation(ctx context.Context, input *models.GiveXpInput) (*models.Discord, error) {
	err := auth.Allow(ctx, []string{"admin", "bot"})
	if err != nil {
		return nil, err
	}

	driver := *global.Get().Driver
	result, err := ogm.GiveXp(ctx, driver, input.DiscordId, input.Amount)
	if err != nil {
		return nil, err
	}

	if result.Records == nil || len(result.Records) == 0 {
		return nil, fmt.Errorf("could not find discord with id %s", input.DiscordId)
	}
	discord, err := misc.MapResult(&models.Discord{}, result)
	if err != nil {
		return nil, err
	}

	levelMaxXp := misc.MaxOutAt(discord.Level*1000, 10000)
	if discord.Xp >= levelMaxXp {
		discord.Level += 1
		discord.Xp = discord.Xp - levelMaxXp
	}

	result, err = ogm.UpdateDiscord(ctx, driver, discord)
	if err != nil {
		return nil, err
	}
	discord, err = misc.MapResult(&models.Discord{}, result)
	if err != nil {
		return nil, err
	}

	return discord, nil
}
