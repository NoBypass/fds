package repository

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"server/src/graph/generated/models"
)

func FindAccountByName(ctx context.Context, driver neo4j.DriverWithContext, name string) (*neo4j.EagerResult, error) {
	result, err := neo4j.ExecuteQuery(ctx, driver,
		"MATCH (a:Account { name: $name }) RETURN a",
		map[string]any{
			"name": name,
		}, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func CreateAccount(ctx context.Context, driver neo4j.DriverWithContext, account *models.Account) (*neo4j.EagerResult, error) {
	result, err := neo4j.ExecuteQuery(ctx, driver,
		"CREATE (a:Account { name: $name, email: $email, password: $password, role: $role, created_at: $created_at }) RETURN a",
		map[string]any{
			"name":       account.Name,
			"password":   account.Password,
			"created_at": account.CreatedAt,
		}, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}

	return result, nil
}
