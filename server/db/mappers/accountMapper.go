package mappers

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
	"server/db/models"
)

func ResultToAccount(result *neo4j.EagerResult) (*models.Account, error) {
	accountNode, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "a")
	if err != nil {
		return nil, err
	}

	id := result.Records[0].Values[0].(dbtype.Node).Id

	username, err := neo4j.GetProperty[string](accountNode, "username")
	if err != nil {
		return nil, err
	}

	password, err := neo4j.GetProperty[string](accountNode, "password")
	if err != nil {
		return nil, err
	}

	joinedAt, err := neo4j.GetProperty[string](accountNode, "joined_at")
	if err != nil {
		return nil, err
	}

	return &models.Account{
		ID:       id,
		Username: username,
		Password: password,
		JoinedAt: joinedAt,
	}, nil
}
