package service

import (
	"encoding/json"
	"github.com/NoBypass/fds/internal/pkg/utils"
	"github.com/NoBypass/surgo"
	"github.com/labstack/gommon/log"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/surrealdb/surrealdb.go"
)

type DatabaseService interface {
	DB(sp opentracing.Span) *surgo.DB
}

type databaseService struct {
	service
	db *surgo.DB
}

func NewDatabaseService(envCfg *utils.Config) DatabaseService {
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

	svc := databaseService{db: &surgo.DB{
		DB: &database{
			DB: db.DB.(*surrealdb.DB),
		},
	}}
	svc.SetName("SurrealDB")
	svc.db.DB.(*database).tracer = svc.tracer

	return svc
}

func (c databaseService) DB(sp opentracing.Span) *surgo.DB {
	return &surgo.DB{
		DB: &database{
			DB:     c.db.DB.(*database).DB,
			tracer: c.db.DB.(*database).tracer,
			sp:     sp,
		},
	}
}

type database struct {
	*surrealdb.DB
	sp     opentracing.Span
	tracer opentracing.Tracer
}

func (qa *database) Query(sql string, vars any) (any, error) {
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
