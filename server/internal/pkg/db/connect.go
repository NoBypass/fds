package db

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/redis/go-redis/v9"
	"server/internal/pkg/misc"
)

func Connect(env misc.ENV) (neo4j.DriverWithContext, *redis.Client) {
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

	err = driver.VerifyConnectivity(context.Background())
	if err != nil {
		panic(err)
	}

	_, err = cache.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	return driver, cache
}
