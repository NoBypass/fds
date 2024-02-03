package repository

import (
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/surrealdb/surrealdb.go"
	"strings"
)

type HypixelRepository interface {
	Repository
	Create(member *model.HypixelPlayer) error
}

type hypixelRepository struct {
	*surrealdb.DB
}

func NewHypixelRepository(db *surrealdb.DB) HypixelRepository {
	return &hypixelRepository{
		db,
	}
}

func (r *hypixelRepository) Create(member *model.HypixelPlayer) error {
	query := strings.Replace(`CREATE hypixel_player:[$uuid, $date] CONTENT {
		"uuid"": $uuid,
		"date"": $date
	};`, "\n", " ", -1)
	_, err := r.DB.Query(query, map[string]interface{}{
		"uuid": member.UUID,
		"date": member.Date,
	})
	return err
}
