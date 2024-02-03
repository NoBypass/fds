package model

type MojangProfile struct {
	Date string
	UUID string `json:"id"`
	Name string `json:"name"`
}
