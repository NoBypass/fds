package controller

import "github.com/NoBypass/fds/internal/domain"

type (
	inputName struct {
		Name string `param:"name"`
	}

	inputPwd struct {
		Pwd string `json:"pwd"`
	}

	inputVerify struct {
		DiscordName string `param:"discord_name"`
		DiscordID   string `param:"discord_id"`
		IGN         string `param:"ign"`
	}

	inputRevoke struct {
		DiscordID string `param:"discord_id"`
	}

	inputXP struct {
		DiscordID string `json:"discord_id"`
		Amount    int    `json:"amount"`
	}

	// TODO: replace any with actual types
	scrimsPlayer struct {
		Player *domain.ScrimsPlayerData `json:"player"`
		Totals *domain.ScrimsMode       `json:"totals"`
	}
)
