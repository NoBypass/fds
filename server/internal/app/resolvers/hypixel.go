package resolvers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"server/internal/pkg/generated/models"
	"server/pkg/ogm"
)

type NewPlayer struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func PlayerQuery(ctx context.Context, input *models.PlayerInput) (*models.Player, error) {
	players := ogm.New[models.Player](ctx)
	player, _ := players.Find(&models.Player{
		Name: input.Name,
	})
	fmt.Printf("%+v\n", player)
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

	newPlayerInput := NewPlayer{}
	if err := json.NewDecoder(response.Body).Decode(&newPlayerInput); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	return players.Create(&models.Player{
		Uuid: newPlayerInput.ID,
		Name: newPlayerInput.Name,
	})
}
