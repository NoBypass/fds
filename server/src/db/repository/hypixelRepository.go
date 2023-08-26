package repository

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
	"net/http"
	"server/src/db/mappers"
	"server/src/db/models"
)

func FindPlayerByName(ctx context.Context, driver neo4j.DriverWithContext, name string) (*models.Player, error) {
	result, err := neo4j.ExecuteQuery(ctx, driver,
		"MATCH (p:Player) WHERE toLower(p.name) = toLower($name) RETURN p",
		map[string]any{
			"name": name,
		}, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}

	if result.Records != nil || len(result.Records) > 0 {
		return mappers.ResultToPlayer(result)
	}

	var url = "https://api.mojang.com/users/profiles/minecraft/" + name
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error making the API request: %v", err)
	}
	if response.StatusCode == 404 {
		return nil, errors.New("player not found")
	}
	defer response.Body.Close()

	var playerDto models.PlayerDto
	if err := json.NewDecoder(response.Body).Decode(&playerDto); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	result, err = neo4j.ExecuteQuery(ctx, driver,
		"CREATE (p:Player { name: $name, uuid: $uuid }) RETURN p",
		map[string]any{
			"name": playerDto.Name,
			"uuid": playerDto.ID,
		}, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}

	return mappers.ResultToPlayer(result)
}
