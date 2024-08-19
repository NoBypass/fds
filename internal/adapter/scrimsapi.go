package adapter

import (
	"context"
	"fmt"
	"github.com/NoBypass/fds/internal/common"
	"github.com/NoBypass/fds/internal/domain"
	"github.com/NoBypass/mincache"
	"time"
)

type ScrimsAPI struct {
	api *common.ExternalClient
}

func NewScrimsAPI(cache *mincache.Cache) *ScrimsAPI {
	return &ScrimsAPI{
		api: common.NewExternalClient(cache, "https://api.scrims.network/v1", "Scrims API"),
	}
}

func (s *ScrimsAPI) PlayerByName(ctx context.Context, name string) (*domain.ScrimsPlayerData, error) {
	return s.player(ctx, "user?username="+name)
}

//func (s *ScrimsAPI) PlayerByUUID(ctx context.Context, uuid string) (*domain.ScrimsPlayerData, error) {
//	return s.player(ctx, "user?uuid="+uuid)
//}
//
//func (s *ScrimsAPI) PlayerByDiscordID(ctx context.Context, id string) (*domain.ScrimsPlayerData, error) {
//	return s.player(ctx, "user?discord_id="+id)
//}

func (s *ScrimsAPI) player(ctx context.Context, url string) (*domain.ScrimsPlayerData, error) {
	// TODO handle `octoberWins` field (example: bdamja)
	player := new(ScrimsPlayerAPIResponse)
	_, err := s.api.Request(ctx, url, time.Minute*5, &player)
	if err != nil {
		return nil, fmt.Errorf("scrims: %w", err)
	} else if player.Data == nil {
		return nil, fmt.Errorf("scrims data for player not found")
	}
	return player.Data, err
}

type ScrimsPlayerAPIResponse struct {
	Success bool                     `json:"success"`
	Data    *domain.ScrimsPlayerData `json:"user_data"`
}
