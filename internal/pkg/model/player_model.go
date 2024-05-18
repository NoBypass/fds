package model

type Player struct {
	Name        string
	UUID        string
	DisplayName string        `db:"display_name"`
	ScrimsData  *ScrimsPlayer `db:"scrims_data"`
}

type ScrimsPlayer struct {
	Data *ScrimsPlayerData `json:"user_data"`
	Date string            `json:"date"`
	UUID string            `json:"uuid"`
}
