package conf

import (
	"fmt"
	"github.com/surrealdb/surrealdb.go"
	"server/internal/core/repository"
)

func (c *Config) ConnectDB() *repository.Repository {
	db, err := surrealdb.New(fmt.Sprintf("ws://%s:%d/rpc", c.Database.Host, c.Database.Port))
	if err != nil {
		panic(err)
	}

	if _, err = db.Signin(map[string]interface{}{
		"NS":   c.Database.Namespace,
		"DB":   c.Database.Name,
		"user": c.Database.User,
		"pass": c.Database.Password,
	}); err != nil {
		panic(err)
	}

	if _, err = db.Use(c.Database.Namespace, c.Database.Name); err != nil {
		panic(err)
	}

	return repository.New(db)
}
