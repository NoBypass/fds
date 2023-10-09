package ogm

import (
	"github.com/mitchellh/mapstructure"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"reflect"
)

// Query will execute a query on the database with the
// parameters passed in. It will return the result of
// the query, or an error if one occurred.
func (ogm *OGM) Query(query string, params map[string]any) ([]*neo4j.Record, error) {
	res, err := neo4j.ExecuteQuery(ogm.ctx, ogm.driver, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}
	return res.Records, nil
}

// Map will map the result of a query to the object
// passed in. It will return the object, or an error
// if one occurred.
func Map[T any](obj *T, records []*neo4j.Record, variable string) (*T, error) {
	idx := 0
	for i, val := range (*records[0]).Keys {
		if val == variable {
			idx = i
			break
		}
	}
	specific := (*records[0]).Values[idx].(neo4j.Node).Props
	objVal := *obj
	inputType := reflect.TypeOf(objVal)
	newMap := make(map[string]any, len(specific))

	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		jsonTag := field.Tag.Get("json")
		newMap[field.Name] = specific[jsonTag]
	}

	err := mapstructure.Decode(newMap, &objVal)
	if err != nil {
		return nil, err
	}

	obj = &objVal
	return obj, nil
}
