package ogm

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// OGM is a struct which contains the driver
// and context for the database. With it
// you can interact with the database with
// type safety.
type OGM struct {
	driver neo4j.DriverWithContext
	ctx    context.Context
}

// New returns a new instance of the OGM struct
// which can be used to interact with the database
// with type safety.
func New(ctx context.Context, driver neo4j.DriverWithContext) *OGM {
	return &OGM{
		driver: driver,
		ctx:    ctx,
	}
}
