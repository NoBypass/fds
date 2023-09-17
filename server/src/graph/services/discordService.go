package services

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"net/http"
	"server/src/api/handlers"
	"server/src/graph/generated/models"
	"server/src/repository"
	"server/src/utils"
)

func CreateDiscordMutation(ctx context.Context, input *models.CreateDiscordInput) (*models.Discord, error) {
	result, err := repository.CreateDiscord(ctx, ctx.Value("driver").(neo4j.DriverWithContext), &models.Discord{
		DiscordId: input.DiscordId,
		Name:      input.Name,
	})
	if err != nil {
		return nil, err
	}

	return utils.MapResult(&models.Discord{}, result, "d")
}

func DiscordQuery(ctx context.Context, input *models.DiscordInput) (*models.Discord, error) {
	result, err := repository.FindDiscordByDiscordId(ctx, ctx.Value("driver").(neo4j.DriverWithContext), input.DiscordId)
	if err != nil {
		return nil, err
	}

	return utils.MapResult(&models.Discord{}, result, "d")
}

func GiveXpMutation(ctx context.Context, input *models.GiveXpInput) (*models.Discord, error) {
	claims, err := utils.ParseJWT(ctx, ctx.Value("token").(string))
	if err != nil {
		return nil, err
	}

	if claims.Role != "bot" && claims.Role != "admin" {
		return nil, handlers.NewHttpError(ctx, http.StatusUnauthorized, "you do not have permission to give xp")
	}

	result, err := repository.GiveXp(ctx, ctx.Value("driver").(neo4j.DriverWithContext), input.DiscordId, input.Amount)
	if err != nil {
		return nil, err
	}

	handlers.CheckIfFound(ctx, result, "could not find discord with id "+input.DiscordId)
	discord, err := utils.MapResult(&models.Discord{}, result, "d")
	if err != nil {
		return nil, err
	}

	levelMaxXp := utils.MaxOutAt(discord.Level*1000, 10000)
	if discord.Xp >= levelMaxXp {
		discord.Level += 1
		discord.Xp = discord.Xp - levelMaxXp
	}

	result, err = repository.UpdateDiscord(ctx, ctx.Value("driver").(neo4j.DriverWithContext), discord)
	if err != nil {
		return nil, err
	}
	discord, err = utils.MapResult(&models.Discord{}, result, "d")
	if err != nil {
		return nil, err
	}

	return discord, nil
}
