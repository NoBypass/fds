package models

// Nodes

type Account struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	JoinedAt int64  `json:"joined_at"`
}

type DiscordAccount struct {
	ID          string `json:"id"`
	DiscordID   string `json:"discord_id"`
	Name        string `json:"name"`
	Level       int64  `json:"level"`
	Xp          int64  `json:"xp"`
	Streak      int64  `json:"streak"`
	LastDailyAt int64  `json:"last_daily_at"`
}

// Edges

type IsLinkedTo struct {
	ID       string `json:"id"`
	LinkedAt int64  `json:"linked_at"`
}
