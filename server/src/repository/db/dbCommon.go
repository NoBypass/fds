package db

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type DB[T any] struct {
	driver neo4j.DriverWithContext
	ctx    context.Context
}

func New[T any](ctx context.Context) *DB[T] {
	return &DB[T]{
		driver: ctx.Value("driver").(neo4j.DriverWithContext),
		ctx:    ctx,
	}
}
