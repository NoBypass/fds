package adapter

import (
	"context"
	"encoding/json"
	"github.com/NoBypass/fds/internal/common/env"
	"github.com/NoBypass/fds/internal/common/trace"
	"github.com/NoBypass/surgo"
	"github.com/labstack/gommon/log"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/surrealdb/surrealdb.go"
)

type Database struct {
	trace.Tracable
	db *surgo.DB
}

type database struct {
	*surrealdb.DB
	trace.Tracable
	ctx context.Context
}

func ConnectSurreal(cfg *env.Env) Database {
	db, err := surgo.Connect(
		cfg.SurrealHost,
		surgo.User(cfg.SurrealUser),
		surgo.Password(cfg.SurrealPwd),
		surgo.Database(cfg.SurrealDB),
		surgo.Namespace(cfg.SurrealNamespace),
	)
	if err != nil {
		log.Fatalf("couldn't connect to SurrealDB: %s", err)
	}

	svc := Database{
		Tracable: trace.NewTracable("SurrealDB"),
		db: &surgo.DB{
			DB: &database{
				DB: db.DB.(*surrealdb.DB),
			},
		},
	}

	return svc
}

func (qa *database) Query(sql string, vars any) (any, error) {
	var sp opentracing.Span
	if qa.ctx != nil {
		sp, _ = qa.Tracable.StartSpan(qa.ctx, sql)
		defer sp.Finish()

		sp.LogKV("sql", sql)
		j, err := json.Marshal(vars)
		if err != nil {
			ext.LogError(sp, err)
		} else {
			sp.LogKV("vars", string(j))
		}
	}

	resp, err := qa.DB.Query(sql, vars)
	if err != nil && sp != nil {
		ext.LogError(sp, err)
	} else if sp != nil {
		j, err := json.Marshal(resp)
		if err != nil {
			ext.LogError(sp, err)
		} else {
			sp.LogKV("response", string(j))
		}
	}

	return resp, err
}

func (c Database) DB(ctx context.Context) *surgo.DB {
	return &surgo.DB{
		DB: &database{
			DB:       c.db.DB.(*database).DB,
			Tracable: c.db.DB.(*database).Tracable,
			ctx:      ctx,
		},
	}
}
