package service

import (
	"github.com/NoBypass/fds/internal/backend/database"
	"github.com/NoBypass/fds/internal/external"
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/NoBypass/mincache"
	"github.com/NoBypass/surgo"
	"github.com/opentracing/opentracing-go"
	"time"
)

type ScrimsService interface {
	Service
	PlayerFromDB(name string) <-chan *model.ScrimsPlayer
	PlayerFromAPI(name string) <-chan *model.ScrimsPlayer
}

type scrimsService struct {
	service
	apiClient *external.ScrimsAPIClient
	database.Client
}

func NewScrimsService(db database.Client, cache *mincache.Cache) ScrimsService {
	return &scrimsService{
		apiClient: external.NewScrimsAPIClient(cache),
		Client:    db,
	}
}

func (s *scrimsService) PlayerFromDB(name string) <-chan *model.ScrimsPlayer {
	out := make(chan *model.ScrimsPlayer)

	s.Pipeline(func(start func() opentracing.Span) error {
		sp := start()
		defer sp.Finish()

		var player model.ScrimsPlayer
		err := s.DB(sp).Scan(&player, "SELECT ONLY * FROM scrims_player:$", surgo.Range{
			surgo.ID{name, time.Unix(0, 0).Format(time.RFC3339)},
			surgo.ID{name, time.Now().Format(time.RFC3339)},
		})
		if err != nil {
			return err
		}

		out <- &player
		close(out)
		return nil
	}, s.PlayerFromDB)

	return out
}

func (s *scrimsService) PlayerFromAPI(name string) <-chan *model.ScrimsPlayer {
	out := make(chan *model.ScrimsPlayer)

	s.Pipeline(func(start func() opentracing.Span) error {
		sp := start()
		defer sp.Finish()

		player, err := s.apiClient.Player(name, sp)
		if err != nil {
			return err
		}

		out <- player
		close(out)
		return nil
	}, s.PlayerFromAPI)

	return out
}
