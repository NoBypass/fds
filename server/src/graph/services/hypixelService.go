package services

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
	"net/http"
	"server/src/graph/generated/models"
	"server/src/repository"
)

func PlayerQuery(ctx context.Context, input *models.PlayerInput) (*models.Player, error) {
	result, err := repository.FindPlayerByName(ctx, ctx.Value("driver").(neo4j.DriverWithContext), input.Name)
	if err != nil {
		return nil, err
	}

	if result.Records != nil || len(result.Records) > 0 {
		return models.ResultToPlayer(result)
	}

	var url = "https://api.mojang.com/users/profiles/minecraft/" + input.Name
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error making the API request: %v", err)
	}
	if response.StatusCode == 404 {
		return nil, errors.New("player not found")
	}
	defer response.Body.Close()

	var newPlayerInput repository.NewPlayer
	if err := json.NewDecoder(response.Body).Decode(&newPlayerInput); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	result, err = repository.CreatePlayer(ctx, ctx.Value("driver").(neo4j.DriverWithContext), &newPlayerInput)
	if err != nil {
		return nil, err
	}

	return models.ResultToPlayer(result)
}
