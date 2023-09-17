package db

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"reflect"
	"server/src/utils"
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

func (db *DB[T]) common(entity *T, prefix string) (*T, error) {
	values, err := utils.StructToMap(entity)
	if err != nil {
		return nil, err
	}

	var args string
	for key := range values {
		args += fmt.Sprintf("%s: $%s,", key, key)
	}

	result, err := neo4j.ExecuteQuery(db.ctx, db.driver,
		fmt.Sprintf("%s (n:%s { %s }) RETURN n", prefix, reflect.TypeOf(*entity).Name(), args[:len(values)-1]), values, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}

	entity, err = utils.MapResult(entity, result, "n")
	return entity, err
}
