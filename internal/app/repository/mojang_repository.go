package repository

import (
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/NoBypass/fds/internal/pkg/surreal_wrap"
	"time"
)

type MojangRepository interface {
	Repository
	Create(member *model.MojangProfile) error
}

type mojangRepository struct {
	repository
}

func NewMojangRepository(db *surreal_wrap.DB) MojangRepository {
	return &mojangRepository{
		newRepository(db),
	}
}

func (r *mojangRepository) Create(member *model.MojangProfile) error {
	now := time.Now()
	date := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Format(time.RFC3339)
	_, err := r.DB.Queryf(`CREATE mojang_profile:["%s", "%s"] CONTENT {
			"name": "%s",
			"date": "%s",
			"uuid": "%s"
		}`, member.Name, date, member.Name, date, member.UUID)
	return err
}
