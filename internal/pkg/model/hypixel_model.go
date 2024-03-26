package model

type HypixelPlayerResponse struct {
	Success bool `json:"success"`
	Player  struct {
		UUID        string `json:"uuid"`
		DisplayName string `json:"displayname"`
		SocialMedia struct {
			Links struct {
				Discord string `json:"DISCORD"`
			} `json:"links"`
		} `json:"socialMedia"`
	} `json:"player"`
}

type HypixelPlayer struct {
	Date string `json:"date"`
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type Unused struct {
	Data string `json:"data"`
}
