package db

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"server/src/utils"
)

type DB[T any] struct {
	driver neo4j.DriverWithContext
	ctx    context.Context
}

func Connect(ctx context.Context) (neo4j.DriverWithContext, error) {
	env := ctx.Value("env").(utils.ENV).DB
	driver, err := neo4j.NewDriverWithContext(env.URI, neo4j.BasicAuth(env.Username, env.Password, ""))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return driver, nil
}

func New[T any](ctx context.Context) *DB[T] {
	return &DB[T]{
		driver: ctx.Value("driver").(neo4j.DriverWithContext),
		ctx:    ctx,
	}
}
