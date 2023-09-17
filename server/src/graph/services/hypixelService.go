package services

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"server/src/graph/generated/models"
	"server/src/repository"
	"server/src/repository/db"
)

func PlayerQuery(ctx context.Context, input *models.PlayerInput) (*models.Player, error) {
	players := db.New[models.Player](ctx)
	player, _ := players.Find(&models.Player{
		Name: input.Name,
	})
	if player != nil {
		return player, nil
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

	newPlayerInput := repository.NewPlayer{}
	if err := json.NewDecoder(response.Body).Decode(&newPlayerInput); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	return players.Create(&models.Player{
		Uuid: newPlayerInput.ID,
		Name: newPlayerInput.Name,
	})
}
