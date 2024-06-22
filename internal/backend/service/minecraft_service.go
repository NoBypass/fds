package service

import (
	"context"
	"github.com/NoBypass/fds/internal/backend/tracing"
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/NoBypass/fds/internal/pkg/utils"
	"github.com/NoBypass/surgo"
	"strings"
	"time"
)

type MinecraftService interface {
	NameToUUID(ctx context.Context, name string) (string, error)
	LocalScrimsStats(ctx context.Context, uuid string, fields []string, date ...time.Time) (map[string]any, error)
	RemoteScrimsStats(ctx context.Context, name string) (*model.ScrimsPlayerData, error)
	TotalScrimsStats(ctx context.Context, raw *model.ScrimsPlayerData) (*model.ScrimsMode, error)
}

type minecraftService struct {
	DatabaseService
	tracing.Tracable

	scrimsSvc ScrimsService
}

func NewMinecraftService(db DatabaseService, scrimsSvc ScrimsService) MinecraftService {
	return &minecraftService{
		DatabaseService: db,
		scrimsSvc:       scrimsSvc,
		Tracable:        tracing.NewTracable(),
	}
}

func (s *minecraftService) NameToUUID(ctx context.Context, name string) (string, error) {
	sp, ctx := s.StartSpan(ctx, s.NameToUUID)
	defer sp.Finish()

	var uuid string
	err := s.DB(sp).Scan(&uuid, "RETURN (SELECT uuid FROM ONLY player:$).uuid", surgo.ID{name})
	if err != nil {
		return "", err
	} else if uuid != "" {
		return uuid, nil
	}

	player, err := s.RemoteScrimsStats(ctx, name)
	if err != nil {
		return "", err
	}

	return player.UUID, nil
}

func (s *minecraftService) LocalScrimsStats(ctx context.Context, uuid string, fields []string, date ...time.Time) (map[string]any, error) {
	return nil, nil
}

func (s *minecraftService) RemoteScrimsStats(ctx context.Context, name string) (*model.ScrimsPlayerData, error) {
	sp, ctx := s.StartSpan(ctx, s.RemoteScrimsStats)
	defer sp.Finish()

	player, err := s.scrimsSvc.PlayerByName(ctx, name)
	if err != nil {
		return nil, err
	}

	_, err = s.DB(sp).Exec(`
	LET $rec = r'player:$';
	LET $player = (SELECT * FROM ONLY $rec);
	
	IF $player = NONE {
		LET $new = (CREATE ONLY $rec CONTENT {
			display_name: $displayName,
			name: $name,
			uuid: $uuid
		});
		CREATE scrims_player:[$new.uuid, $today] CONTENT {
			data: $data,
			date: $today,
			uuid: $uuid
		};
		UPDATE $new SET scrims_data=scrims_player:[$new.uuid, $today];
	} ELSE {
		IF $player.scrims_data.date == $today {
			UPDATE scrims_player:[$player.uuid, $today] MERGE {
				data: $data
			};
		} ELSE {
			CREATE scrims_player:[$player.uuid, $today] CONTENT {
				data: $data,
				date: $today,
				uuid: $uuid
			};
			UPDATE $player SET scrims_data=scrims_player:[$player.uuid, $today];
		};
	};
	`, surgo.ID{strings.ToLower(name)}, map[string]any{
		"name":        strings.ToLower(player.Data.Username),
		"today":       utils.Today(),
		"displayName": player.Data.Username,
		"uuid":        player.Data.UUID,
		"data":        player.Data,
	})
	if err != nil {
		return nil, err
	}

	return player.Data, nil
}

func (s *minecraftService) TotalScrimsStats(ctx context.Context, raw *model.ScrimsPlayerData) (*model.ScrimsMode, error) {
	sp, ctx := s.StartSpan(ctx, s.TotalScrimsStats)
	defer sp.Finish()

	bridge := raw.Stats.Bridge
	modes := []map[string]model.ScrimsMode{raw.Stats.Tow.Duel, bridge.Casual, bridge.Duel, bridge.Ranked, bridge.Private}

	var total model.ScrimsMode
	for _, mode := range modes {
		for _, m := range mode {
			total.Wins += m.Wins
			total.Games += m.Games
			total.Kills += m.Kills
			total.Goals += m.Goals
			total.Draws += m.Draws
			total.Losses += m.Losses
			total.Deaths += m.Deaths
			total.ArrowsHit += m.ArrowsHit
			total.HitsGiven += m.HitsGiven
			total.HitsTaken += m.HitsTaken
			total.ArrowsShot += m.ArrowsShot
			total.HitsBlocked += m.HitsBlocked
			total.BlocksPlaced += m.BlocksPlaced
			total.BlocksBroken += m.BlocksBroken
			total.GapplesEaten += m.GapplesEaten
			total.PlayerCausedDeaths += m.PlayerCausedDeaths
			total.YLevelSum += m.YLevelSum
			total.DamageDealt += m.DamageDealt
			total.IGT += m.IGT
		}
	}

	return &total, nil
}
