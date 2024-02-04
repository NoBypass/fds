package conf

import (
	"fmt"
	"github.com/NoBypass/fds/internal/pkg/surreal_wrap"
	"github.com/surrealdb/surrealdb.go"
)

func (c *Config) ConnectDB() *surreal_wrap.DB {
	db, err := surrealdb.New(fmt.Sprintf("wss://%s/rpc", c.DBHost))
	if err != nil {
		db, err = surrealdb.New(fmt.Sprintf("ws://%s/rpc", c.DBHost))
		if err != nil {
			panic(err)
		}
	}

	if _, err = db.Signin(map[string]interface{}{
		"NS":   c.DBNamespace,
		"DB":   c.DBName,
		"user": c.DBUser,
		"pass": c.DBPwd,
	}); err != nil {
		panic(err)
	}

	if _, err = db.Use(c.DBNamespace, c.DBName); err != nil {
		panic(err)
	}

	return surreal_wrap.Wrap(db)
}
