package db

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"net/http"
	"reflect"
	"server/src/api/handlers"
	"server/src/utils"
)

// Find will find a node in the database by the values
// of the struct passed in. It will return the node
// found in the database, or an error if one occurred.
func (db *DB[T]) Find(entity *T) (*T, error) {
	values := utils.StructToMap(entity)

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
		res := db.ctx.Value("res").(*handlers.Responder)
		res.Status(http.StatusNotFound)
		return nil, res.AddError(fmt.Errorf(name+" not found"), handlers.NODE_NOT_FOUND, []string{"db.go"})
	}

	entity, err = utils.MapResult(entity, result, "n")
	return entity, err
}

// Create will create a node in the database by the struct
// passed in. It will return the node created in the database,
// or an error if one occurred.
func (db *DB[T]) Create(entity *T) (*T, error) {
	values := utils.StructToMap(entity)

	var args string
	for key := range values {
		args += fmt.Sprintf(" n.%s = $%s AND", key, key)
	}

	result, err := neo4j.ExecuteQuery(db.ctx, db.driver,
		fmt.Sprintf("MATCH (n:%s) WHERE%sRETURN n", reflect.TypeOf(*entity).Name(), args[:len(args)-3]), values, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}

	entity, err = utils.MapResult(entity, result, "n")
	return entity, err
}
