package ogm

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"server/internal/app/global"
)

// OGM is a struct which contains the driver
// and context for the database. With it
// you can interact with the database with
// type safety.
type OGM[T any] struct {
	driver neo4j.DriverWithContext
	ctx    context.Context
}

// New returns a new instance of the OGM struct
// which can be used to interact with the database
// with type safety.
func New[T any](ctx context.Context) *OGM[T] {
	return &OGM[T]{
		driver: *global.Get().Driver,
		ctx:    ctx,
	}
}
