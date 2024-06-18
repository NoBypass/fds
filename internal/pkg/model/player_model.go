package model

import "time"

type Player struct {
	Name        string
	UUID        string
	DisplayName string        `db:"display_name"`
	ScrimsData  *ScrimsPlayer `db:"scrims_data"`
}

type ScrimsPlayer struct {
	Date time.Time         `json:"date"`
	Data *ScrimsPlayerData `json:"data"`
}
