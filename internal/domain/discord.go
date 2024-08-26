package domain

import (
	"math"
	"math/rand"
	"time"
)

type Leaderboard []struct {
	DiscordID string `json:"discord_id"`
	Xp        int    `json:"xp"`
	Level     int    `json:"level"`
}

type DiscordMember struct {
	DiscordID   string    `json:"discord_id"`
	LastDailyAt time.Time `json:"last_daily_at"`
	Xp          int       `json:"xp"`
	Level       int       `json:"level"`
	Streak      int       `json:"streak"`
}

func (m *DiscordMember) CanClaimDaily() bool {
	return m.LastDailyAt.After(time.Now().Truncate(time.Hour * 24))
}

func (m *DiscordMember) ClaimDaily() {
	m.AddXP(int(rand.Float64() * 500 * (1 + float64(m.Streak)*0.1)))
	m.LastDailyAt = time.Now()
	m.Streak++
}

func (m *DiscordMember) AddXP(xp int) {
	m.Xp += xp
	needed := m.GetNeededXP()
	if m.Xp >= needed {
		m.Level++
		m.Xp = m.Xp - needed
	}
}

func (m *DiscordMember) GetNeededXP() int {
	if m.Level < 10 {
		return int(math.Pow(float64(m.Level), 2) * 100)
	}
	return 10000
}
