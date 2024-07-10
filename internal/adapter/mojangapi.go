package adapter

import (
	"context"
	"github.com/NoBypass/fds/internal/common"
	"github.com/NoBypass/fds/internal/domain"
	"github.com/NoBypass/mincache"
	"time"
)

type MojangAPI struct {
	cache *mincache.Cache
	api   *common.ExternalClient
}

func NewMojangAPI(cache *mincache.Cache) *MojangAPI {
	return &MojangAPI{
		cache: cache,
		api:   common.NewExternalClient(cache, "https://api.mojang.com", "Mojang API"),
	}
}

func (m *MojangAPI) ProfileByName(ctx context.Context, name string) (*domain.MojangProfile, error) {
	player := new(domain.MojangProfile)
	_, err := m.api.Request(ctx, "/users/profiles/minecraft/"+name, 5*time.Minute, player)
	if err != nil {
		return nil, err
	}
	return player, nil
}
