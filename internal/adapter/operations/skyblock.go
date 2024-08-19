package operations

import (
	"github.com/NoBypass/fds/internal/adapter"
	"github.com/NoBypass/fds/internal/common/trace"
)

type SkyblockRepository struct {
	trace.Tracable
	adapter.Database
}

func NewSkyblockRepository(db adapter.Database) *SkyblockRepository {
	return &SkyblockRepository{
		Tracable: trace.NewTracable(),
		Database: db,
	}
}
