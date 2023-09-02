package repository

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type NewPlayer struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func FindPlayerByName(ctx context.Context, driver neo4j.DriverWithContext, name string) (*neo4j.EagerResult, error) {
	result, err := neo4j.ExecuteQuery(ctx, driver,
		"MATCH (p:Player) WHERE toLower(p.name) = toLower($name) RETURN p",
		map[string]any{
			"name": name,
		}, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func CreatePlayer(ctx context.Context, driver neo4j.DriverWithContext, player *NewPlayer) (*neo4j.EagerResult, error) {
	result, err := neo4j.ExecuteQuery(ctx, driver,
		"CREATE (p:Player { name: $name, uuid: $uuid }) RETURN p",
		map[string]any{
			"name": player.Name,
			"uuid": player.ID,
		}, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}

	return result, nil
}
