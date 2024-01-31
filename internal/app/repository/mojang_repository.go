package repository

import (
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/surrealdb/surrealdb.go"
)

type MojangRepository interface {
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
	member.ID = "mojang_profile:" + member.ID
	_, err := surrealdb.SmartMarshal(r.DB.Create, member)
	return err
}
