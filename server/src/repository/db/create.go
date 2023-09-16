package db

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"golang.org/x/exp/maps"
	"reflect"
	"server/src/utils"
)

func (db *DB[T]) Create(entity *T) (*T, error) {
	values := make(map[string]any)
	input := reflect.TypeOf(*entity)
	for i := 0; i < input.NumField(); i++ {
		field := input.Field(i)
		values[utils.ConvertCamelToSnake(field.Name)] = reflect.ValueOf(entity).Elem().FieldByName(field.Name).Interface()
	}

	var args string
	for key := range values {
		if key == maps.Keys(values)[len(values)-1] {
			args += fmt.Sprintf("%s: $%s", key, key)
			continue
		}
		args += fmt.Sprintf("%s: $%s, ", key, key)
	}

	result, err := neo4j.ExecuteQuery(db.ctx, db.driver,
		fmt.Sprintf("CREATE (n:%s { %s }) RETURN n", input.Name(), args), values, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}

	entity, err = utils.MapResult(entity, result, "n")
	return entity, err
}
