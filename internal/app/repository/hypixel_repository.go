package repository

import (
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/NoBypass/fds/internal/pkg/surreal_wrap"
	"time"
)

type HypixelRepository interface {
	Repository
	Create(member *model.HypixelPlayer) error
}

type hypixelRepository struct {
	repository
}

func NewHypixelRepository(db *surreal_wrap.DB) HypixelRepository {
	return &hypixelRepository{
		newRepository(db),
	}
}

func (r *hypixelRepository) Create(member *model.HypixelPlayer) error {
	now := time.Now()
	member.Date = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Format(time.RFC3339)
	_, err := r.DB.Queryf(`CREATE hypixel_player:["%s", "%s"] CONTENT {
		"uuid": "%s",
		"date": "%s"
	}`, member.UUID, member.Date, member.UUID, member.Date)
	return err
}
