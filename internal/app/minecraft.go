package app

import (
	"context"
	"github.com/NoBypass/fds/internal/common/trace"
	"github.com/NoBypass/fds/internal/domain"
	"time"
)

type MinecraftRepository interface {
	GetPlayer(ctx context.Context, name string) (*domain.ScrimsPlayerData, time.Time, error)
	UpsertScrimsPlayer(ctx context.Context, player *domain.ScrimsPlayerData) error
}

type ScrimsService interface {
	PlayerByName(ctx context.Context, name string) (*domain.ScrimsPlayerData, error)
}

type MojangService interface {
	ProfileByName(ctx context.Context, name string) (*domain.MojangProfile, error)
}

type MinecraftUseCase struct {
	trace.Tracable

	repo   MinecraftRepository
	scrims ScrimsService
	mojang MojangService
}

func NewMinecraftUseCase(
	repo MinecraftRepository,
	scrims ScrimsService,
	mojang MojangService) *MinecraftUseCase {
	return &MinecraftUseCase{
		Tracable: trace.NewTracable(),
		repo:     repo,
		scrims:   scrims,
		mojang:   mojang,
	}
}

func (uc *MinecraftUseCase) GetScrimsPlayer(ctx context.Context, name string) (*domain.ScrimsPlayerData, error) {
	sp, ctx := uc.StartSpan(ctx, uc.GetScrimsPlayer)
	defer sp.Finish()

	player, err := uc.scrims.PlayerByName(ctx, name)
	if err != nil {
		return nil, err
	}

	err = uc.repo.UpsertScrimsPlayer(ctx, player)
	if err != nil {
		return nil, err
	}

	return player, nil
}

func (uc *MinecraftUseCase) GetMojangProfile(ctx context.Context, name string) (*domain.MojangProfile, error) {
	sp, ctx := uc.StartSpan(ctx, uc.GetMojangProfile)
	defer sp.Finish()

	profile, err := uc.mojang.ProfileByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return profile, nil
}
