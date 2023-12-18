package model

type DiscordMember struct {
	ID     string  `json:"id"`
	Nick   string  `json:"nick"`
	XP     float64 `json:"xp"`
	Level  int     `json:"level"`
	Streak int     `json:"streak"`
}
