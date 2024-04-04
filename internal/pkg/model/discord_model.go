package model

import (
	"math"
	"time"
)

type (
	DiscordMember struct {
		DiscordID   string  `json:"discord_id" db:"discord_id"`
		Name        string  `json:"name"`
		Nick        string  `json:"nick"`
		XP          float64 `json:"xp"`
		LastDailyAt string  `json:"last_daily_at" db:"last_daily_at"`
		Level       int     `json:"level"`
		Streak      int     `json:"streak"`
	}

	DiscordVerifyRequest struct {
		ID   string `json:"id"`
		Nick string `json:"nick"`
		Name string `json:"name"`
	}

	DiscordVerifyResponse struct {
		Actual string `json:"actual"`
	}

	DiscordDailyResponse struct {
		XP        float64 `json:"xp"`
		Level     int     `json:"level"`
		Levelup   bool    `json:"levelup"`
		Needed    float64 `json:"needed"`
		Gained    float64 `json:"gained"`
		Streak    int     `json:"streak"`
		WithBonus float64 `json:"with_bonus"`
	}

	DiscordID struct {
		ID   string `json:"id"`
		Date string `json:"date"`
	}

	DiscordLeaderboardEntry struct {
		DiscordID string  `json:"discord_id" db:"discord_id"`
		Level     int     `json:"level"`
		XP        float64 `json:"xp"`
	}

	DiscordBotLoginResponse struct {
		Token string `json:"token"`
	}

	DiscordBotLoginRequest struct {
		Pwd string `json:"pwd" query:"pwd"`
	}
)

func (d *DiscordMember) AddXP(xp float64) {
	d.XP += xp
	needed := d.GetNeededXP()
	if d.XP >= needed {
		d.Level++
		d.XP = d.XP - needed
	}
}

func (d *DiscordMember) GetNeededXP() float64 {
	if d.Level < 10 {
		return math.Pow(float64(d.Level), 2) * 100
	}
	return 10000
}

func (d *DiscordMember) CanClaimDaily() bool {
	timestamp, _ := time.Parse(time.RFC3339, d.LastDailyAt)
	return timestamp.Add(24*time.Hour).Sub(time.Now()) < 0
}
