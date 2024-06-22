package model

import "time"

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

type ScrimsPlayerAPIResponse struct {
	Success bool              `json:"success"`
	Data    *ScrimsPlayerData `json:"user_data"`
}

type ScrimsPlayerTimes struct {
	Date       time.Time `json:"date"`
	LastLogin  int       `json:"lastLogin" db:"last_login"`
	LastLogout int       `json:"lastLogout" db:"last_logout"`
	Playtime   int       `json:"playtime"`
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
