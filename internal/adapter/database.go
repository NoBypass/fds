package adapter

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/NoBypass/fds/internal/common/env"
	"github.com/NoBypass/fds/internal/common/trace"
	"github.com/NoBypass/surgo/v2"
	"github.com/labstack/gommon/log"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

type Database struct {
	trace.Tracable
	db *surgo.DB
}

type database struct {
	trace.Tracable
	db  surgo.DBConn
	ctx context.Context
}

func ConnectSurreal(cfg *env.Env) Database {
	db, err := surgo.Connect(cfg.SurrealHost, &surgo.Credentials{
		Username:  cfg.SurrealUser,
		Password:  cfg.SurrealPwd,
		Database:  cfg.SurrealDB,
		Namespace: cfg.SurrealNamespace,
	})
	if err != nil {
		log.Fatalf("couldn't connect to SurrealDB: %s", errors.Unwrap(err))
	}

	svc := Database{
		Tracable: trace.NewTracable("SurrealDB"),
		db:       db,
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

	resp, err := qa.db.Query(sql, vars)
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

func (qa *database) Close() error {
	return qa.db.Close()
}

func (c Database) DB(ctx context.Context) *surgo.DB {
	return &surgo.DB{
		Conn: &database{
			Tracable: c.Tracable,
			db:       c.db.Conn,
			ctx:      ctx,
		},
	}
}
