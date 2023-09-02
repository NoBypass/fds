package models

import "github.com/neo4j/neo4j-go-driver/v5/neo4j"

// Code automatically generated; DO NOT EDIT.

type PlayerInput struct {
	Name string `json:"name"`
}

type Player struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}

func ResultToPlayer(result *neo4j.EagerResult) (*Player, error) {
	r, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "p")
	if err != nil {
		return nil, err
	}

	UUID, err := neo4j.GetProperty[string](r, "uuid")
	if err != nil {
		return nil, err
	}

	name, err := neo4j.GetProperty[string](r, "name")
	if err != nil {
		return nil, err
	}
	return &Player{
		Uuid: UUID,
		Name: name,
	}, nil
}
