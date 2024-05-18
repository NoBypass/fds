package model

type ScrimsPlayerData struct {
	Cages      []string `json:"cages"`
	LastLogin  int      `json:"lastLogin"`
	LastLogout int      `json:"lastLogout"`
	Playtime   int      `json:"playtime"`
	Username   string   `json:"username"`
	DiscordID  string   `json:"discordId"`
	Ranked     struct {
		S1 struct {
			Elo    int `json:"elo"`
			Games  int `json:"games"`
			Losses int `json:"losses"`
			Wins   int `json:"wins"`
		} `json:"s1"`
	} `json:"ranked"`
	Stats struct {
		Bridge struct {
			Casual  map[string]mode `json:"casual"`
			Duel    map[string]mode `json:"duel"`
			Ranked  map[string]mode `json:"ranked"`
			Private map[string]mode `json:"private"`
			Overall overall         `json:"overall"`
		} `json:"bridge"`
		Tow struct {
			Duel map[string]mode `json:"duel"`
		} `json:"tow"`
		Overall overall `json:"overall"`
	} `json:"stats"`
}

type ScrimsPlayerResponse struct {
	Success bool              `json:"success"`
	Data    *ScrimsPlayerData `json:"user_data"`
}

type overall struct {
	Winstreak         int `json:"winstreak"`
	DailyWinstreak    int `json:"dailyWinstreak"`
	LifetimeWinstreak int `json:"lifetimeWinstreak"`
}

type mode struct {
	Wins                int     `json:"wins"`
	Games               int     `json:"games"`
	Kills               int     `json:"kills"`
	Goals               int     `json:"goals"`
	Draws               int     `json:"draws"`
	Losses              int     `json:"losses"`
	Deaths              int     `json:"deaths"`
	ArrowsHit           int     `json:"arrowsHit"`
	HitsGiven           int     `json:"hitsGiven"`
	HitsTaken           int     `json:"hitsTaken"`
	ArrowsShot          int     `json:"arrowsShot"`
	HitsBlocked         int     `json:"hitsBlocked"`
	BlocksPlaced        int     `json:"blocksPlaced"`
	BlocksBroken        int     `json:"blocksBroken"`
	GapplesEaten        int     `json:"gapplesEaten"`
	PlayerCausedDeaths  int     `json:"playerCausedDeaths"`
	YLevelSum           float64 `json:"yLevelSum"`
	DamageDealt         float64 `json:"damageDealt"`
	SecondsSpentPlaying float64 `json:"secondsSpentPlaying"`
}
