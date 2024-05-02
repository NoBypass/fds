package model

type ScrimsPlayer struct {
	Success bool `json:"success"`
	Data    struct {
		Cages      []string `json:"cages"`
		LastLogin  string   `json:"lastLogin"`
		LastLogout string   `json:"lastLogout"`
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
	} `json:"user_data"`
}

type overall struct {
	Winstreak         int `json:"winstreak"`
	DailyWinstreak    int `json:"dailyWinstreak"`
	LifetimeWinstreak int `json:"lifetimeWinstreak"`
}

type mode struct {
	Wins                int     `json:"wins"`
	Losses              int     `json:"losses"`
	Games               int     `json:"games"`
	Kills               int     `json:"kills"`
	Deaths              int     `json:"deaths"`
	PlayerCausedDeaths  int     `json:"playerCausedDeaths"`
	Goals               int     `json:"goals"`
	BlocksPlaced        int     `json:"blocksPlaced"`
	BlocksBroken        int     `json:"blocksBroken"`
	DamageDealt         int     `json:"damageDealt"`
	GapplesEaten        int     `json:"gapplesEaten"`
	ArrowsShot          int     `json:"arrowsShot"`
	ArrowsHit           int     `json:"arrowsHit"`
	HitsGiven           int     `json:"hitsGiven"`
	HitsTaken           int     `json:"hitsTaken"`
	HitsBlocked         int     `json:"hitsBlocked"`
	YLevelSum           float64 `json:"yLevelSum"`
	SecondsSpentPlaying int     `json:"secondsSpentPlaying"`
	Draws               int     `json:"draws"`
}
