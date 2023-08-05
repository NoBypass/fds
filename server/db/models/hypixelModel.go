package models

// Nodes

type Player struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

// Dto

type PlayerDto struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}
