package repository

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"server/db/models"
)

func CreateAccount(ctx context.Context, driver neo4j.DriverWithContext, account *models.Account) (*models.Account, error) {
	result, err := neo4j.ExecuteQuery(ctx, driver,
		"CREATE (a:Account { id: $id, username: $username, password: $password, joined_at: $joined_at }) RETURN a",
		map[string]any{
			"id":        account.ID,
			"username":  account.Username,
			"password":  account.Password,
			"joined_at": account.JoinedAt,
		}, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}

	accountNode, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "a")
	if err != nil {
		return nil, err
	}

	id, err := neo4j.GetProperty[string](accountNode, "id")
	if err != nil {
		return nil, err
	}

	username, err := neo4j.GetProperty[string](accountNode, "username")
	if err != nil {
		return nil, err
	}

	password, err := neo4j.GetProperty[string](accountNode, "password")
	if err != nil {
		return nil, err
	}

	joinedAt, err := neo4j.GetProperty[int64](accountNode, "joined_at")
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

func findAccountByName(ctx context.Context, driver neo4j.DriverWithContext, name string) (*models.Account, error) {
	result, err := neo4j.ExecuteQuery(ctx, driver,
		"MATCH (a:Account { username: $username }) RETURN a",
		map[string]any{
			"username": name,
		}, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}

	accountNode, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "a")
	if err != nil {
		return nil, err
	}

	id, err := neo4j.GetProperty[string](accountNode, "id")
	if err != nil {
		return nil, err
	}

	username, err := neo4j.GetProperty[string](accountNode, "username")
	if err != nil {
		return nil, err
	}

	password, err := neo4j.GetProperty[string](accountNode, "password")
	if err != nil {
		return nil, err
	}

	joinedAt, err := neo4j.GetProperty[int64](accountNode, "joined_at")
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
