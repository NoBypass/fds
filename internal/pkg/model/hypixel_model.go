package model

type HypixelPlayerResponse struct {
	Success bool `json:"success"`
	Player  struct {
		Id          string `json:"_id"`
		DisplayName string `json:"displayname"`
		SocialMedia struct {
			Links struct {
				Discord string `json:"DISCORD"`
			} `json:"links"`
		} `json:"socialMedia"`
	} `json:"player"`
}

type HypixelPlayer string
