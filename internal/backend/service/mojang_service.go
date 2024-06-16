package service

import (
	"fmt"
	"github.com/NoBypass/fds/internal/backend/database"
	"github.com/NoBypass/fds/internal/external"
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/NoBypass/mincache"
	"github.com/NoBypass/surgo"
	"github.com/opentracing/opentracing-go/ext"
	"strings"
	"time"
)

type MojangService interface {
	Service
	PersistPlayer(*model.MojangProfile) (*model.Player, error)

	PlayerFromAPI(string) (*model.MojangProfile, error)
	PlayerFromDB(string, ...string) (*model.Player, error)
}

type mojangService struct {
	service
	apiClient *external.MojangAPIClient
	database.Client
}

func NewMojangService(db database.Client, cache *mincache.Cache) MojangService {
	return &mojangService{
		apiClient: external.NewMojangAPIClient(cache),
		Client:    db,
	}
}

func (s *mojangService) PlayerFromDB(name string, fields ...string) (*model.Player, error) {
	end, sp := s.Trace(s.PlayerFromDB)
	defer end()

	var fieldsStr strings.Builder
	fieldsStr.WriteString("*")
	if len(fields) > 0 {
		for i, field := range fields {
			if i == 0 {
				fieldsStr.Reset()
			} else {
				fieldsStr.WriteString(", ")
			}
			fieldsStr.WriteString(field)
		}
	}

	player := new(model.Player)
	err := s.DB(sp).Scan(player,
		fmt.Sprintf("SELECT %s FROM ONLY player:$", fieldsStr.String()),
		surgo.ID{strings.ToLower(name)})

	if err != nil {
		ext.LogError(sp, err)
		return nil, err
	}

	return player, nil
}

func (s *mojangService) PlayerFromAPI(name string) (*model.MojangProfile, error) {
	end, sp := s.Trace(s.PlayerFromAPI)
	defer end()

	profile, err := s.apiClient.Player(name, sp)
	if err != nil {
		ext.LogError(sp, err)
		return nil, err
	}

	profile.Date = time.Now().Format(time.DateOnly)
	return profile, nil
}

func (s *mojangService) PersistPlayer(player *model.MojangProfile) (*model.Player, error) {
	end, sp := s.Trace(s.PersistPlayer)
	defer end()

	playerID := surgo.ID{strings.ToLower(player.Name)}

	// TODO only update if already exists

	newPlayer := new(model.Player)
	err := s.DB(sp).Scan(newPlayer, `
	CREATE ONLY player:$ CONTENT {
		name: $1,
		uuid: $uuid,
		date: <datetime>$date,
		display_name: $name,
	}`, playerID, *player, strings.ToLower(player.Name))
	if err != nil {
		ext.LogError(sp, err)
		return nil, err
	}

	return newPlayer, nil
}
