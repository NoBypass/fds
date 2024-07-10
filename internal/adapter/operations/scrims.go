package operations

import (
	"context"
	"github.com/NoBypass/fds/internal/adapter"
	"github.com/NoBypass/fds/internal/common"
	"github.com/NoBypass/fds/internal/common/trace"
	"github.com/NoBypass/fds/internal/domain"
	"github.com/NoBypass/surgo"
	"strings"
	"time"
)

type ScrimsRepository struct {
	trace.Tracable
	adapter.Database
}

type scrimsPlayer struct {
	Data *domain.ScrimsPlayerData `json:"data"`
	Date time.Time                `json:"date"`
}

func NewScrimsRepository(db adapter.Database) *ScrimsRepository {
	return &ScrimsRepository{
		Tracable: trace.NewTracable(),
		Database: db,
	}
}

func (r *ScrimsRepository) GetPlayer(ctx context.Context, name string) (*domain.ScrimsPlayerData, time.Time, error) {
	var player scrimsPlayer
	err := r.DB(ctx).Scan(&player, "SELECT * FROM player:$", surgo.ID{strings.ToLower(name)})
	return player.Data, player.Date, err
}

func (r *ScrimsRepository) UpsertPlayer(ctx context.Context, player *domain.ScrimsPlayerData) error {
	_, err := r.DB(ctx).Exec(`
		LET $new = (UPSERT ONLY player:$ CONTENT {
			display_name: $displayName,
			name: $name,
			uuid: $uuid
		});
		UPSERT scrims_player:[$new.uuid, $today] CONTENT {
			data: $data,
			date: $today,
			uuid: $uuid
		};
		UPDATE $new SET scrims_data=scrims_player:[$new.uuid, $today];
	`, surgo.ID{strings.ToLower(player.Username)}, map[string]any{
		"name":         strings.ToLower(player.Username),
		"today":        common.Today(),
		"display_name": player.Username,
		"uuid":         player.UUID,
		"data":         player,
	})
	return err
}
