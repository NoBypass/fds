package ogm

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"reflect"
	"server/internal/pkg/misc"
)

// Find will find a node in the database by the values
// of the struct passed in. It will return the node
// found in the database, or an error if one occurred.
func (db *OGM[T]) Find(entity *T) (*T, error) {
	values := misc.StructToMap(entity)

	var args string
	for key := range values {
		args += fmt.Sprintf(" toLower(n.%s) = toLower($%s) AND", key, key)
	}

	name := reflect.TypeOf(*entity).Name()
	result, err := neo4j.ExecuteQuery(db.ctx, db.driver,
		fmt.Sprintf("MATCH (n:%s) WHERE%sRETURN n", name, args[:len(args)-3]), values, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}

	if result.Records == nil || len(result.Records) == 0 {
		return nil, fmt.Errorf(name + " not found")
	}

	entity, err = misc.MapResult(entity, result)
	return entity, err
}

// Create will create a node in the database by the struct
// passed in. It will return the node created in the database,
// or an error if one occurred.
func (db *OGM[T]) Create(entity *T) (*T, error) {
	values := misc.StructToMap(entity)

	var args string
	for key := range values {
		args += fmt.Sprintf(" n.%s = $%s AND", key, key)
	}

	result, err := neo4j.ExecuteQuery(db.ctx, db.driver,
		fmt.Sprintf("MATCH (n:%s) WHERE%sRETURN n", reflect.TypeOf(*entity).Name(), args[:len(args)-3]), values, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}

	entity, err = misc.MapResult(entity, result)
	return entity, err
}
