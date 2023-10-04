package db

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/redis/go-redis/v9"
	"server/internal/pkg/misc"
	"server/src/api/handlers/logger"
)

func Connect(ctx context.Context) (neo4j.DriverWithContext, *redis.Client, error) {
	env := ctx.Value("env").(misc.ENV)
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
