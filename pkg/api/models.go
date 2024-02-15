package api

type DiscordVerifyRequest struct {
	ID   string `json:"id"`
	Nick string `json:"nick"`
	Name string `json:"name"`
}

type DiscordVerifyResponse struct {
	Actual string `json:"actual"`
}

type DiscordMemberResponse struct {
	DiscordID   string  `json:"discord_id"`
	Name        string  `json:"name"`
	Nick        string  `json:"nick"`
	XP          float64 `json:"xp"`
	LastDailyAt string  `json:"last_daily_at"`
	Level       int     `json:"level"`
	Streak      int     `json:"streak"`
}

type DiscordBotLoginRequest struct {
	Pwd string `json:"pwd" query:"pwd"`
}

type DiscordBotLoginResponse struct {
	Token string `json:"token"`
}
