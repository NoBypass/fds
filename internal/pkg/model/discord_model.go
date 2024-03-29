package model

import (
	"math"
	"time"
)

type DiscordMember struct {
	DiscordID   string  `json:"discord_id"`
	Name        string  `json:"name"`
	Nick        string  `json:"nick"`
	XP          float64 `json:"xp"`
	LastDailyAt string  `json:"last_daily_at"`
	Level       int     `json:"level"`
	Streak      int     `json:"streak"`
}

type DiscordVerifyInput struct {
	ID   string `json:"id"`
	Nick string `json:"nick"`
	Name string `json:"name"`
}

type DiscordBotLoginInput struct {
	Pwd string `json:"pwd" query:"pwd"`
}

type DiscordVerifyResponse struct {
	Name string `json:"name"`
}

type DiscordDailyResponse struct {
	XP        float64 `json:"xp"`
	Level     int     `json:"level"`
	Levelup   bool    `json:"levelup"`
	Needed    float64 `json:"needed"`
	Gained    float64 `json:"gained"`
	Streak    int     `json:"streak"`
	WithBonus float64 `json:"with_bonus"`
}

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
