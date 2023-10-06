package global

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/redis/go-redis/v9"
	"server/internal/pkg/db"
	"server/internal/pkg/misc"
	"sync"
)

type Context struct {
	Env    *misc.ENV
	Driver *neo4j.DriverWithContext
	Cache  *redis.Client
}

var store *Context
var once sync.Once

func Get() *Context {
	once.Do(func() {
		ctx := context.Background()
		env := misc.FetchEnv()
		driver, cache := db.Connect(ctx)
		store = &Context{
			Env:    &env,
			Driver: &driver,
			Cache:  cache,
		}
	})

	return store
}
