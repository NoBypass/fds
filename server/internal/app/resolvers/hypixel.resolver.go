package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"server/internal/pkg/generated/models"
	"server/pkg/ogm"
)

// Player is the resolver for the player field.
func (r *queryResolver) Player(ctx context.Context, name string) (*models.Player, error) {
	pre := ogm.WithPreload(ctx, r.OGM, &models.Player{})
	player, err := pre.Find(map[string]any{"name": name}, "WHERE toLower(p.name) = toLower($name)")
	if player != nil {
		return player, err
	}

	response, err := http.Get("https://api.mojang.com/users/profiles/minecraft/" + name)
	if err != nil || response.StatusCode == 404 {
		return nil, fmt.Errorf("could not find %s on mojang api", name)
	}
	defer response.Body.Close()
	newPlayer := NewPlayer{}
	if err := json.NewDecoder(response.Body).Decode(&newPlayer); err != nil {
		return nil, err
	}

	records, err := r.OGM.Query("CREATE (p:Player {name: $name, uuid: $uuid})", map[string]any{
		"name": newPlayer.Name,
		"uuid": newPlayer.ID,
	})
	if err != nil {
		return nil, err
	}

	return ogm.Map(player, records, "p")
}