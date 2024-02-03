package repository

import (
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/surrealdb/surrealdb.go"
	"strings"
	"time"
)

type MojangRepository interface {
	Repository
	Create(member *model.MojangProfile) error
}

type mojangRepository struct {
	*surrealdb.DB
}

func NewMojangRepository(db *surrealdb.DB) MojangRepository {
	return &mojangRepository{
		db,
	}
}

func (r *mojangRepository) Create(member *model.MojangProfile) error {
	now := time.Now()
	query := strings.Replace(`CREATE mojang_profile:[$name, $date] CONTENT {
			"name": $name,
			"date": $date,
			"uuid": $uuid
		};`, "\n", " ", -1)
	_, err := r.DB.Query(query, map[string]interface{}{
		"name": member.Name,
		"date": time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Format(time.RFC3339),
		"uuid": member.UUID,
	})
	return err
}
