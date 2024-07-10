package operations

import (
	"context"
	"github.com/NoBypass/fds/internal/domain"
)

type PlayerRepository interface {
	CreatePlayer(ctx context.Context, player *domain.Player) error
}

type playerRepository struct {
}
