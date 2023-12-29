package model

import "math"

type DiscordMember struct {
	Nick   string  `json:"nick"`
	XP     float64 `json:"xp"`
	Level  int     `json:"level"`
	Streak int     `json:"streak"`
}

type DiscordSignupInput struct {
	ID   string `json:"id"`
	Nick string `json:"nick"`
}

type DiscordDailyInput struct {
	ID string `json:"id"`
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
