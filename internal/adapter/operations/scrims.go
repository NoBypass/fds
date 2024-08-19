package operations

import (
	"context"
	"fmt"
	"github.com/NoBypass/fds/internal/adapter"
	"github.com/NoBypass/fds/internal/common"
	"github.com/NoBypass/fds/internal/common/trace"
	"github.com/NoBypass/fds/internal/domain"
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
	err := r.DB(ctx).Scan(&player, "SELECT * FROM $player", map[string]any{
		"player": fmt.Sprintf("player:%s", strings.ToLower(name)),
	})
	return player.Data, player.Date, err
}

func (r *ScrimsRepository) UpsertScrimsPlayer(ctx context.Context, player *domain.ScrimsPlayerData) error {
	_, err := r.DB(ctx).Query(`
		LET $new = (UPSERT ONLY $player CONTENT {
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
	`, map[string]any{
		"player":       fmt.Sprintf("player:%s", strings.ToLower(player.Username)),
		"name":         strings.ToLower(player.Username),
		"today":        common.Today(),
		"display_name": player.Username,
		"uuid":         player.UUID,
		"data":         player,
	})
	return err
}

func (r *ScrimsRepository) UpsertPlayer(ctx context.Context, player *domain.MojangProfile) error {
	_, err := r.DB(ctx).Query(`
		UPSERT ONLY $player CONTENT {
			uuid: $uuid,
			name: $name,
			display_name: $name,
		};
	`, map[string]any{
		"player":       fmt.Sprintf("player:%s", strings.ToLower(player.Name)),
		"uuid":         player.UUID,
		"name":         strings.ToLower(player.Name),
		"display_name": player.Name,
	})

	return err
}
