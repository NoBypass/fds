package app

import (
	"context"
	"github.com/NoBypass/fds/internal/common/trace"
	"github.com/NoBypass/fds/internal/domain"
)

type SkyblockRepository interface {
	RelateAuctioneer(ctx context.Context, auctioneer string) error
}

type SkyblockService interface {
	Auctions(ctx context.Context, page int) (*domain.HypixelAuctionResponse, error)
}

type SkyblockUseCase struct {
	trace.Tracable

	repo SkyblockRepository
	svc  SkyblockService
}

func NewSkyblockUseCase(
	repo SkyblockRepository,
	svc SkyblockService) *SkyblockUseCase {
	return &SkyblockUseCase{
		Tracable: trace.NewTracable(),
		repo:     repo,
		svc:      svc,
	}
}

func (uc *SkyblockUseCase) GetAuctions(ctx context.Context) (*domain.HypixelAuctionResponse, error) {
	// TODO: Implement this
	return nil, nil
}
