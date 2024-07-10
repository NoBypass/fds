package domain

type ScrimsPlayerData struct {
	UUID       string   `json:"_id"`
	Cages      []string `json:"cages"`
	LastLogin  int      `json:"lastLogin" db:"lastLogin"`
	LastLogout int      `json:"lastLogout" db:"lastLogout"`
	Playtime   int      `json:"playtime"`
	Username   string   `json:"username"`
	DiscordID  string   `json:"discordId" db:"discordId"`
	Ranked     map[string]struct {
		Elo    float64 `json:"elo"`
		Games  int     `json:"games"`
		Losses int     `json:"losses"`
		Wins   int     `json:"wins"`
	} `json:"ranked"`
	Stats struct {
		Bridge struct {
			Casual  map[string]ScrimsMode `json:"casual"`
			Duel    map[string]ScrimsMode `json:"duel"`
			Ranked  map[string]ScrimsMode `json:"ranked"`
			Private map[string]ScrimsMode `json:"private"`
			Overall overall               `json:"overall"`
		} `json:"bridge"`
		Tow struct {
			Duel map[string]ScrimsMode `json:"duel"`
		} `json:"tow"`
		Overall overall `json:"overall"`
	} `json:"stats"`
}

type overall struct {
	Winstreak         int `json:"winstreak"`
	DailyWinstreak    int `json:"dailyWinstreak"`
	LifetimeWinstreak int `json:"lifetimeWinstreak"`
}

type ScrimsMode struct {
	Wins               int     `json:"wins"`
	Games              int     `json:"games"`
	Kills              int     `json:"kills"`
	Goals              int     `json:"goals"`
	Draws              int     `json:"draws"`
	Losses             int     `json:"losses"`
	Deaths             int     `json:"deaths"`
	ArrowsHit          int     `json:"arrowsHit" db:"arrowsHit"`
	HitsGiven          int     `json:"hitsGiven" db:"hitsGiven"`
	HitsTaken          int     `json:"hitsTaken" db:"hitsTaken"`
	ArrowsShot         int     `json:"arrowsShot" db:"arrowsShot"`
	HitsBlocked        int     `json:"hitsBlocked" db:"hitsBlocked"`
	BlocksPlaced       int     `json:"blocksPlaced" db:"blocksPlaced"`
	BlocksBroken       int     `json:"blocksBroken" db:"blocksBroken"`
	GapplesEaten       int     `json:"gapplesEaten" db:"gapplesEaten"`
	PlayerCausedDeaths int     `json:"playerCausedDeaths" db:"playerCausedDeaths"`
	YLevelSum          float64 `json:"yLevelSum" db:"yLevelSum"`
	DamageDealt        float64 `json:"damageDealt" db:"damageDealt"`
	IGT                float64 `json:"secondsSpentPlaying" db:"secondsSpentPlaying"`
}

func (s *ScrimsPlayerData) Totals(raw *ScrimsPlayerData) (*ScrimsMode, error) {
	bridge := raw.Stats.Bridge
	modes := []map[string]ScrimsMode{raw.Stats.Tow.Duel, bridge.Casual, bridge.Duel, bridge.Ranked, bridge.Private}

	var total ScrimsMode
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
