package service

import (
	"context"
	"github.com/NoBypass/fds/internal/backend/tracing"
	"github.com/NoBypass/fds/internal/external"
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/NoBypass/mincache"
	"time"
)

type ScrimsService interface {
	PlayerByName(context.Context, string) (*model.ScrimsPlayerAPIResponse, error)
	PlayerByUUID(context.Context, string) (*model.ScrimsPlayerAPIResponse, error)
	PlayerByDiscordID(context.Context, string) (*model.ScrimsPlayerAPIResponse, error)
	// Leaderboard(string) ...

	AllPlayerTimes(string) ([]*model.ScrimsPlayerTimes, error)
	PlayerByDate(string, time.Time) (*model.ScrimsPlayer, error)

	PersistScrimsPlayer(*model.ScrimsPlayerAPIResponse, *model.Player) (*model.ScrimsPlayerData, error)
	PersistPlayer(*model.ScrimsPlayerAPIResponse) (*model.Player, error)

	PlayerFromDB(string) (*model.ScrimsPlayerAPIResponse, error)
	PlayerFromAPI(string) (*model.ScrimsPlayerAPIResponse, error)
}

type scrimsService struct {
	DatabaseService
	tracing.Tracable

	api *external.Client
}

func NewScrimsService(db DatabaseService, cache *mincache.Cache) ScrimsService {
	return &scrimsService{
		api:             external.NewClient(cache, "https://api.scrims.network/v1", "Scrims API"),
		DatabaseService: db,
		Tracable:        tracing.NewTracable(),
	}
}

func (s *scrimsService) PlayerByName(ctx context.Context, name string) (*model.ScrimsPlayerAPIResponse, error) {
	sp, ctx := s.StartSpan(ctx, s.PlayerByName)
	defer sp.Finish()

	player := new(model.ScrimsPlayerAPIResponse)
	_, err := s.api.Request(ctx, "user?username="+name, time.Minute*5, player)
	return player, err
}

func (s *scrimsService) PlayerByUUID(ctx context.Context, uuid string) (*model.ScrimsPlayerAPIResponse, error) {
	sp, ctx := s.StartSpan(ctx, s.PlayerByUUID)
	defer sp.Finish()

	player := new(model.ScrimsPlayerAPIResponse)
	_, err := s.api.Request(ctx, "user?uuid="+uuid, time.Minute*5, &player)
	return player, err
}

func (s *scrimsService) PlayerByDiscordID(ctx context.Context, id string) (*model.ScrimsPlayerAPIResponse, error) {
	sp, ctx := s.StartSpan(ctx, s.PlayerByDiscordID)
	defer sp.Finish()

	player := new(model.ScrimsPlayerAPIResponse)
	_, err := s.api.Request(ctx, "user?discord_id="+id, time.Minute*5, &player)
	return player, err
}

func (s *scrimsService) AllPlayerTimes(name string) ([]*model.ScrimsPlayerTimes, error) {
	//end, sp := s.Trace(s.AllPlayerTimes)
	//defer end()
	//
	//entries := make([]*model.ScrimsPlayerTimes, 0)
	//err := s.DB(sp).Scan(&entries, `
	//LET $uuid = <uuid>(SELECT uuid FROM ONLY player:$).uuid;
	//SELECT date, data.playtime AS playtime, data.lastLogin AS last_login, data.lastLogout AS last_logout FROM scrims_player:[$uuid, NONE]..[$uuid, time::now()];
	//`, surgo.ID{strings.ToLower(name)})
	//if err != nil {
	//	return nil, err
	//}
	//
	//return entries, nil
	return nil, nil
}

func (s *scrimsService) PlayerByDate(name string, date time.Time) (*model.ScrimsPlayer, error) {
	//end, sp := s.Trace(s.PlayerByDate)
	//defer end()
	//
	//entry := new(model.ScrimsPlayer)
	//err := s.DB(sp).Scan(entry, `
	//SELECT * FROM ONLY scrims_player:[<uuid>(SELECT uuid FROM ONLY player:$).uuid, $1]`, surgo.ID{strings.ToLower(name)}, date)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return entry, nil
	return nil, nil
}

func (s *scrimsService) PlayerFromDB(name string) (*model.ScrimsPlayerAPIResponse, error) {
	//end, sp := s.Trace(s.PlayerFromDB)
	//defer end()
	//
	//player := new(model.ScrimsPlayerAPIResponse)
	//err := s.DB(sp).Scan(player, "SELECT scrims_data FROM ONLY player:$", surgo.ID{strings.ToLower(name)})
	//if err != nil {
	//	return nil, err
	//}
	//
	//return player, nil
	return nil, nil
}

func (s *scrimsService) PlayerFromAPI(name string) (*model.ScrimsPlayerAPIResponse, error) {
	//end, sp := s.Trace(s.PlayerFromAPI)
	//defer end()
	//
	//player, err := s.api.Player(name, sp)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return player, nil
	return nil, nil
}

func (s *scrimsService) PersistScrimsPlayer(playerResp *model.ScrimsPlayerAPIResponse, dbPlayer *model.Player) (*model.ScrimsPlayerData, error) {
	//end, sp := s.Trace(s.PersistScrimsPlayer)
	//defer end()
	//
	//today, _ := time.Parse(time.DateOnly, time.Now().Format(time.DateOnly))
	//playerID := surgo.ID{strings.ToLower(playerResp.Data.Username)}
	//scrimsPlayerID := surgo.ID{dbPlayer.UUID, today}
	//
	//var err error
	//if dbPlayer.ScrimsData == nil || !dbPlayer.ScrimsData.Date.Equal(today) {
	//	_, err = s.DB(sp).Exec("CREATE scrims_player:$ CONTENT {data: $1, date: $2}",
	//		scrimsPlayerID,
	//		playerResp.Data,
	//		today)
	//} else {
	//	_, err = s.DB(sp).Exec("UPDATE scrims_player:$ SET scrims_data=$1", scrimsPlayerID, playerResp.Data)
	//}
	//
	//if err != nil {
	//	ext.LogError(sp, err)
	//	return nil, err
	//}
	//
	//if !dbPlayer.ScrimsData.Date.Equal(today) {
	//	_, err = s.DB(sp).Exec("UPDATE player:$ SET scrims_data=scrims_player:$", playerID, scrimsPlayerID)
	//	if err != nil {
	//		return nil, err
	//	}
	////}
	//
	//return playerResp.Data, nil
	return nil, nil
}

func (s *scrimsService) PersistPlayer(player *model.ScrimsPlayerAPIResponse) (*model.Player, error) {
	//end, sp := s.Trace(s.PersistPlayer)
	//defer end()
	//
	//name := strings.ToLower(player.Data.Username)
	//playerID := surgo.ID{name}
	//
	//newPlayer := new(model.Player)
	//err := s.DB(sp).Scan(newPlayer, `
	//CREATE ONLY player:$ CONTENT {
	//	name: $1,
	//	uuid: <uuid>$2,
	//	display_name: $3,
	//}`, playerID, name, player.Data.UUID, player.Data.Username)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return newPlayer, nil
	return nil, nil
}
