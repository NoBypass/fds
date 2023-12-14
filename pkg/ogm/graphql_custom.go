package ogm

import (
	"context"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"reflect"
	"server/internal/pkg/misc"
	"strings"
)

type Constructor[T any] struct {
	Queries map[string]string
	Fields  map[string][]string
	NameMap map[string]string
	Root    *T
	ctx     context.Context
	*OGM
}

func WithPreload[T any](ctx context.Context, ogm *OGM, root *T) *Constructor[T] {
	pre := misc.GetPreloads(ctx, strings.ToLower(reflect.TypeOf(*root).Name()))
	queries, fields, nameMap := generateCypherQuery(pre)

	return &Constructor[T]{
		Queries: queries,
		NameMap: nameMap,
		Fields:  fields,
		Root:    root,
		ctx:     ctx,
		OGM:     ogm,
	}
}

func (c *Constructor[T]) Find(args map[string]any, extra string) (*T, error) {
	query := fmt.Sprintf("MATCH %s", misc.JoinStrMap(c.Queries, " MATCH "))
	query = query[:len(query)-len("MATCH ")]
	query += extra + " RETURN " + misc.JoinStrArrMap(c.Fields, ", ")
	records, err := c.OGM.Query(query[:len(query)-2], args)
	if err != nil {
		return nil, err
	}

	if records == nil {
		return nil, fmt.Errorf("no records found")
	}

	return c.Map(records)
}

func (c *Constructor[T]) Map(records []*neo4j.Record) (*T, error) {
	m := make(map[string]any)
	for qk, _ := range c.Queries {
		subelements := strings.Split(qk, ".")
		working := m
		for i, subelement := range subelements {
			noSnake := misc.DeSnake(subelement)
			if i != 0 {
				working[noSnake] = make(map[string]any)
				working, _ = working[noSnake].(map[string]any)
			}

			for j, key := range records[0].Keys {
				s := strings.Split(key, ".")
				if s[len(s)-2] == c.NameMap[subelement] {
					working[misc.DeSnake(s[len(s)-1])] = records[0].Values[j]
				}
			}
		}
	}

	err := mapstructure.Decode(m, c.Root)
	if err != nil {
		return nil, err
	}
	return c.Root, nil
}
