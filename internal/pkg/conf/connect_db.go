package conf

import (
	"fmt"
	"github.com/surrealdb/surrealdb.go"
)

func (c *Config) ConnectDB() *surrealdb.DB {
	db, err := surrealdb.New(fmt.Sprintf("wss://%s/rpc", c.DBHost))
	if err != nil {
		panic(err)
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

	return db
}
