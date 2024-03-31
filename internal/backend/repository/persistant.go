package repository

import (
	"encoding/json"
	"github.com/NoBypass/surgo"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/surrealdb/surrealdb.go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"os"
)

var (
	db *surgo.DB
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
		j, err := json.Marshal(vars)
		if err != nil {
			ext.LogError(sp, err)
		} else {
			sp.LogKV("response", string(j))
		}
	}

	return resp, err
}

func (qa *Agent) Close() {
	qa.DB.Close()
}

func init() {
	db = surgo.MustConnect(
		os.Getenv("db_host"),
		surgo.Password(os.Getenv("db_pwd")),
		surgo.User(os.Getenv("db_user")),
		surgo.Database(os.Getenv("db_name")),
		surgo.Namespace(os.Getenv("db_namespace")),
	)

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
		panic(err)
	}

	db = &surgo.DB{
		DB: &Agent{
			DB:     db.DB.(*surrealdb.DB),
			tracer: tracer,
		},
	}

	println("✓ Connected to SurrealDB")
}

func DB(sp opentracing.Span) *surgo.DB {
	return &surgo.DB{
		DB: &Agent{
			DB:     db.DB.(*Agent).DB,
			tracer: db.DB.(*Agent).tracer,
			sp:     sp,
		},
	}
}
