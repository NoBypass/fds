package service

import (
	"github.com/NoBypass/fds/internal/backend/database"
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/NoBypass/surgo"
	"github.com/opentracing/opentracing-go/ext"
	"strings"
)

type PlayerService interface {
	Service
	FromDB(string) (*model.Player, error)
}

type playerService struct {
	service
	database.Client
}

func NewPlayerService(db database.Client) PlayerService {
	return &playerService{
		Client: db,
	}
}

func (s *playerService) FromDB(name string) (*model.Player, error) {
	end, sp := s.Trace(s.FromDB)
	defer end()

	player := new(model.Player)
	err := s.DB(sp).Scan(player, "SELECT * FROM ONLY player:$", surgo.ID{strings.ToLower(name)})
	if err != nil {
		ext.LogError(sp, err)
		return nil, err
	}

	return player, nil
}
