package app

import (
	"context"
	"github.com/NoBypass/fds/internal/common/trace"
	"github.com/NoBypass/fds/internal/domain"
	"time"
)

type MinecraftRepository interface {
	GetPlayer(ctx context.Context, name string) (*domain.ScrimsPlayerData, time.Time, error)
	UpsertPlayer(ctx context.Context, player *domain.ScrimsPlayerData) error
}

type MinecraftService interface {
	PlayerByName(ctx context.Context, name string) (*domain.ScrimsPlayerData, error)
}

type MinecraftUseCase struct {
	trace.Tracable

	repo    MinecraftRepository
	service MinecraftService
}

func NewMinecraftUseCase(repo MinecraftRepository, service MinecraftService) *MinecraftUseCase {
	return &MinecraftUseCase{
		Tracable: trace.NewTracable(),
		service:  service,
		repo:     repo,
	}
}

func (uc *MinecraftUseCase) GetScrimsPlayer(ctx context.Context, name string) (*domain.ScrimsPlayerData, error) {
	sp, ctx := uc.StartSpan(ctx, uc.GetScrimsPlayer)
	defer sp.Finish()

	player, err := uc.service.PlayerByName(ctx, name)
	if err != nil {
		return nil, err
	}

	err = uc.repo.UpsertPlayer(ctx, player)
	if err != nil {
		return nil, err
	}

	return player, nil
}
