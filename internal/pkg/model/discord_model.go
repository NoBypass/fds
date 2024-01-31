package model

import (
	"math"
	"time"
)

type DiscordMember struct {
	ID          string  `json:"id"`
	Nick        string  `json:"nick"`
	XP          float64 `json:"xp"`
	LastDailyAt int64   `json:"last_daily_at"`
	Level       int     `json:"level"`
	Streak      int     `json:"streak"`
}

type DiscordVerifyInput struct {
	ID   string `json:"id"`
	Nick string `json:"nick"`
}

type DiscordBotLoginInput struct {
	Pwd string `json:"pwd" query:"pwd"`
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
	return d.LastDailyAt+24*60*60*1000 < time.Now().UnixMilli()
}
