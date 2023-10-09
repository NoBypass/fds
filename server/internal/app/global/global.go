package global

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/redis/go-redis/v9"
	"server/internal/pkg/db"
	"server/internal/pkg/misc"
	"server/pkg/ogm"
	"sync"
)

type Context struct {
	Env    *misc.ENV
	Driver neo4j.DriverWithContext
	Cache  *redis.Client
	DB     *ogm.OGM
}

var store *Context
var once sync.Once

func Get() *Context {
	once.Do(func() {
		env := misc.FetchEnv()
		driver, cache := db.Connect(env)
		store = &Context{
			Env:    &env,
			Driver: driver,
			Cache:  cache,
			DB:     ogm.New(context.Background(), driver),
		}
	})

	return store
}
