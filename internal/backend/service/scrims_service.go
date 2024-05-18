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
		ext.LogError(sp, err)
		return nil, err
	}

	return player, nil
}

func (s *scrimsService) PlayerFromAPI(name string) (*model.ScrimsPlayerResponse, error) {
	end, sp := s.Trace(s.PlayerFromAPI)
	defer end()

	player, err := s.apiClient.Player(name, sp)
	if err != nil {
		ext.LogError(sp, err)
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
	givenDate, err := time.Parse(time.RFC3339, dbPlayer.ScrimsData.Date)
	if err != nil {
		ext.LogError(sp, err)
		return nil, err
	}

	if dbPlayer.ScrimsData == nil || givenDate.Equal(today) {
		_, err = s.DB(sp).Exec("CREATE scrims_player:$ CONTENT {data: $1}",
			scrimsPlayerID,
			playerResp.Data)
	} else {
		_, err = s.DB(sp).Exec("UPDATE scrims_player:$ SET scrims_data=$1", scrimsPlayerID, playerResp.Data)
	}

	if err != nil {
		ext.LogError(sp, err)
		return nil, err
	}

	_, err = s.DB(sp).Exec("UPDATE player:$ SET scrims_data=scrims_player:$", playerID, scrimsPlayerID)
	if err != nil {
		ext.LogError(sp, err)
		return nil, err
	}

	return playerResp.Data, nil
}
