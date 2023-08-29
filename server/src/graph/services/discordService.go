package services

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"server/src/graph/generated"
	"server/src/repository"
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
