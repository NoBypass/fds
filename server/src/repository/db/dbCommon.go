package db

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/redis/go-redis/v9"
	"server/src/api/handlers/logger"
	"server/src/utils"
)

type DB[T any] struct {
	driver neo4j.DriverWithContext
	ctx    context.Context
}

func Connect(ctx context.Context) (neo4j.DriverWithContext, *redis.Client, error) {
	env := ctx.Value("env").(utils.ENV)
	driver, err := neo4j.NewDriverWithContext("neo4j://"+env.Persistent.URI,
		neo4j.BasicAuth(env.Persistent.Username, env.Persistent.Password, ""),
	)
	if err != nil {
		logger.Error(err)
		return nil, nil, err
	}

	cache := redis.NewClient(&redis.Options{
		Addr:     env.Cache.URI,
		Password: env.Cache.Password,
		DB:       0,
	})

	return driver, cache, nil
}

func New[T any](ctx context.Context) *DB[T] {
	return &DB[T]{
		driver: ctx.Value("driver").(neo4j.DriverWithContext),
		ctx:    ctx,
	}
}
