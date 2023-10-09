package resolvers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"server/internal/app/global"
	"server/internal/pkg/generated/models"
	"server/pkg/ogm"
)

type NewPlayer struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func PlayerQuery(ctx context.Context, input *models.PlayerInput) (*models.Player, error) {
	db := global.Get().DB
	res, _ := db.Query("MATCH (p:Player) WHERE toLower(p.name) = toLower($name) RETURN p", map[string]any{
		"name": input.Name,
	})
	player, _ := ogm.Map(&models.Player{}, res, "p")
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

	res, err = db.Query("CREATE (p:Player { name: $name, uuid: $uuid }) RETURN p",
		map[string]any{
			"name": newPlayerInput.Name,
			"uuid": newPlayerInput.ID,
		})
	if err != nil {
		return nil, err
	}
	return ogm.Map(&models.Player{}, res, "p")
}
