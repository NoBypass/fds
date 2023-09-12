package models

// Code automatically generated; DO NOT EDIT.

type PlayerInput struct {
	Name string `json:"name"`
}

type Player struct {
	Uuid         string        `json:"uuid"`
	Name         string        `json:"name"`
	VerifiedWith *VerifiedWith `json:"verified_with"`
}

type VerifiedWith struct {
	VerifiedAt string   `json:"verified_at"`
	Player     *Player  `json:"player"`
	Discord    *Discord `json:"discord"`
}
