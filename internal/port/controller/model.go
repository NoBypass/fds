package controller

import "github.com/NoBypass/fds/internal/domain"

type (
	inputName struct {
		Name string `param:"name"`
	}

	inputPwd struct {
		Pwd string `json:"pwd"`
	}

	// TODO: replace any with actual types
	scrimsPlayer struct {
		Player *domain.ScrimsPlayerData `json:"player"`
		Totals *domain.ScrimsMode       `json:"totals"`
	}
)
