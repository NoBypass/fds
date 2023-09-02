package services

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"net/http"
	"server/src/api/handlers"
	"server/src/graph/generated"
	"server/src/repository"
	"server/src/utils"
)

func CreateDiscordMutation(ctx context.Context, input *generated.CreateDiscordInput) (*generated.Discord, error) {
	result, err := repository.CreateDiscord(ctx, ctx.Value("driver").(neo4j.DriverWithContext), &generated.Discord{
		DiscordId: input.DiscordId,
		Name:      input.Name,
	})
	if err != nil {
		return nil, err
	}

	return generated.ResultToDiscord(result)
}

func DiscordQuery(ctx context.Context, input *generated.DiscordInput) (*generated.Discord, error) {
	result, err := repository.FindDiscordByDiscordId(ctx, ctx.Value("driver").(neo4j.DriverWithContext), input.DiscordId)
	if err != nil {
		return nil, err
	}

	return generated.ResultToDiscord(result)
}

func GiveXpMutation(ctx context.Context, input *generated.GiveXpInput) (*generated.Discord, error) {
	claims, err := handlers.ParseJWT(ctx.Value("token").(string))
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
	discord, err := generated.ResultToDiscord(result)
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
	discord, err = generated.ResultToDiscord(result)
	if err != nil {
		return nil, err
	}

	return discord, nil
}
