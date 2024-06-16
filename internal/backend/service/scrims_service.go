package service

import (
	"github.com/NoBypass/fds/internal/backend/database"
	"github.com/NoBypass/fds/internal/external"
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/NoBypass/mincache"
	"github.com/NoBypass/surgo"
	"github.com/opentracing/opentracing-go/ext"
	"strings"
	"time"
)

type ScrimsService interface {
	Service
	PersistScrimsPlayer(*model.ScrimsPlayerResponse, *model.Player) (*model.ScrimsPlayerData, error)
	PersistPlayer(*model.ScrimsPlayerResponse) (*model.Player, error)

	PlayerFromDB(string) (*model.ScrimsPlayerResponse, error)
	PlayerFromAPI(string) (*model.ScrimsPlayerResponse, error)
}

type scrimsService struct {
	service
	apiClient *external.ScrimsAPIClient
	database.Client
}

func NewScrimsService(db database.Client, cache *mincache.Cache) ScrimsService {
	return &scrimsService{
		apiClient: external.NewScrimsAPIClient(cache),
		Client:    db,
	}
}

func (s *scrimsService) PlayerFromDB(name string) (*model.ScrimsPlayerResponse, error) {
	end, sp := s.Trace(s.PlayerFromDB)
	defer end()

	player := new(model.ScrimsPlayerResponse)
	err := s.DB(sp).Scan(player, "SELECT scrims_data FROM ONLY player:$", surgo.ID{strings.ToLower(name)})
	if err != nil {
		return nil, err
	}

	return player, nil
}

func (s *scrimsService) PlayerFromAPI(name string) (*model.ScrimsPlayerResponse, error) {
	end, sp := s.Trace(s.PlayerFromAPI)
	defer end()

	player, err := s.apiClient.Player(name, sp)
	if err != nil {
		return nil, err
	}

	return player, nil
}

func (s *scrimsService) PersistScrimsPlayer(playerResp *model.ScrimsPlayerResponse, dbPlayer *model.Player) (*model.ScrimsPlayerData, error) {
	end, sp := s.Trace(s.PersistScrimsPlayer)
	defer end()

	today, _ := time.Parse(time.DateOnly, time.Now().Format(time.DateOnly))
	playerID := surgo.ID{strings.ToLower(playerResp.Data.Username)}
	scrimsPlayerID := surgo.ID{dbPlayer.UUID, today}

	var err error
	if dbPlayer.ScrimsData == nil || !dbPlayer.ScrimsData.Date.Equal(today) {
		_, err = s.DB(sp).Exec("CREATE scrims_player:$ CONTENT {data: $1, date: $2}",
			scrimsPlayerID,
			playerResp.Data,
			today)
	} else {
		_, err = s.DB(sp).Exec("UPDATE scrims_player:$ SET scrims_data=$1", scrimsPlayerID, playerResp.Data)
	}

	if err != nil {
		ext.LogError(sp, err)
		return nil, err
	}

	if !dbPlayer.ScrimsData.Date.Equal(today) {
		_, err = s.DB(sp).Exec("UPDATE player:$ SET scrims_data=scrims_player:$", playerID, scrimsPlayerID)
		if err != nil {
			return nil, err
		}
	}

	return playerResp.Data, nil
}

func (s *scrimsService) PersistPlayer(player *model.ScrimsPlayerResponse) (*model.Player, error) {
	end, sp := s.Trace(s.PersistPlayer)
	defer end()

	name := strings.ToLower(player.Data.Username)
	playerID := surgo.ID{name}

	newPlayer := new(model.Player)
	err := s.DB(sp).Scan(newPlayer, `
	CREATE ONLY player:$ CONTENT {
		name: $1,
		uuid: $2,
		display_name: $3,
	}`, playerID, name, player.Data.UUID, player.Data.Username)
	if err != nil {
		return nil, err
	}

	return newPlayer, nil
}
