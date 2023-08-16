package mappers

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
	"server/src/db/models"
)

func ResultToPlayer(result *neo4j.EagerResult) (*models.Player, error) {
	playerNode, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "p")
	if err != nil {
		return nil, err
	}

	id := result.Records[0].Values[0].(dbtype.Node).Id

	name, err := neo4j.GetProperty[string](playerNode, "name")
	if err != nil {
		return nil, err
	}

	uuid, err := neo4j.GetProperty[string](playerNode, "uuid")
	if err != nil {
		return nil, err
	}

	return &models.Player{
		ID:   id,
		Name: name,
		UUID: uuid,
	}, nil
}
