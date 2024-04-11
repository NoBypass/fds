package database

import (
	"encoding/json"
	"github.com/NoBypass/fds/internal/pkg/utils"
	"github.com/NoBypass/surgo"
	"github.com/labstack/gommon/log"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/surrealdb/surrealdb.go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

type Agent struct {
	*surrealdb.DB
	sp     opentracing.Span
	tracer opentracing.Tracer
}

func (qa *Agent) Query(sql string, vars any) (any, error) {
	var sp opentracing.Span
	if qa.sp != nil {
		sp = qa.tracer.StartSpan("Query", opentracing.ChildOf(qa.sp.Context()))
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

func Connect(envCfg *utils.Config) Client {
	db, err := surgo.Connect(
		envCfg.DBHost,
		surgo.User(envCfg.DBUser),
		surgo.Password(envCfg.DBPwd),
		surgo.Database(envCfg.DBName),
		surgo.Namespace(envCfg.DBNamespace),
	)
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.Configuration{
		ServiceName: "SurrealDB",
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}

	tracer, _, err := cfg.NewTracer()
	if err != nil {
		log.Fatal(err)
	}

	db = &surgo.DB{
		DB: &Agent{
			DB:     db.DB.(*surrealdb.DB),
			tracer: tracer,
		},
	}

	return Client{db}
}

type Client struct {
	db *surgo.DB
}

func (c Client) DB(sp opentracing.Span) *surgo.DB {
	return &surgo.DB{
		DB: &Agent{
			DB:     c.db.DB.(*Agent).DB,
			tracer: c.db.DB.(*Agent).tracer,
			sp:     sp,
		},
	}
}
