package surreal_wrap

import (
	"fmt"
	"github.com/surrealdb/surrealdb.go"
	"golang.org/x/net/context"
	"strings"
)

type DB struct {
	*surrealdb.DB
	ctx context.Context
}

func Wrap(db *surrealdb.DB) *DB {
	return &DB{DB: db}
}

func (d *DB) InjectContext(ctx context.Context) {
	d.ctx = ctx
}

func (d *DB) query(query string, args []interface{}) (interface{}, error) {
	if d.ctx == nil {
		return d.DB.Query(query, args)
	}
	select {
	case <-d.ctx.Done():
		return nil, d.ctx.Err()
	default:
		return d.DB.Query(query, args)
	}
}

func (d *DB) Queryf(format string, args ...interface{}) (interface{}, error) {
	if !strings.HasSuffix(format, ";") {
		format += ";"
	}
	query := fmt.Sprintf(format, args...)
	query = strings.Replace(query, "\n", " ", -1)
	query = strings.Replace(query, "\t", "", -1)
	return d.query(query, nil)
}
