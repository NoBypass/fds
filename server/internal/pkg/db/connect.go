package db

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/redis/go-redis/v9"
	"server/internal/pkg/misc"
)

func Connect(ctx context.Context) (neo4j.DriverWithContext, *redis.Client) {
	env := ctx.Value("env").(misc.ENV)
	driver, err := neo4j.NewDriverWithContext("neo4j://"+env.Persistent.URI,
		neo4j.BasicAuth(env.Persistent.Username, env.Persistent.Password, ""),
	)
	if err != nil {
		panic(err)
	}

	cache := redis.NewClient(&redis.Options{
		Addr:     env.Cache.URI,
		Password: env.Cache.Password,
		DB:       0,
	})

	return driver, cache
}
