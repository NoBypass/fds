package ogm

import (
	"context"
	"fmt"
	"reflect"
	"server/internal/pkg/misc"
	"strings"
	"unicode"
)

type Constructor[T any] struct {
	Selectors []string
	Root      *T
}

func WithPreload[T any](ctx context.Context, root *T) *Constructor[T] {
	pre := misc.GetPreloads(ctx, strings.ToLower(reflect.TypeOf(*root).Name()))
	// [name LINKED_TO LINKED_TO.linkedAt] -> `(p:Player)-[l:LINKED_TO]->()` & `RETURN p.name, l.linkedAt`
	// TODO: implement preloads to query conversion
	cypherQuery, rv := generateCypherQuery(pre)
	fmt.Println(cypherQuery, rv)

	return &Constructor[T]{}
}

func generateCypherQuery(preloads []string) (map[string]string, map[string][]string) {
	queries := make(map[string]string)
	returnVars := make(map[string][]string)

	for _, preload := range preloads {
		split := strings.Split(preload, ".")

		varName := strings.ToLower(split[len(split)-2][:1]) + fmt.Sprintf("%d", len(split)-1)
		varVal := strings.ToLower(split[len(split)-1])
		returnVars[varName] = append(returnVars[varName], varVal)

		str := strings.Join(split[:len(split)-1], ".")
		if _, ok := queries[str]; ok {
			continue
		}

		filter := strings.Join(split[:len(split)-2], ".")
		if _, ok := queries[filter]; ok {
			delete(queries, filter)
		}

		queries[str] = ""
	}

	for k, _ := range queries {
		queries[k] = generator(k, &returnVars)
	}
	return queries, returnVars
}

func generator(preload string, returnVars *map[string][]string) string {
	split := strings.Split(preload, ".")
	query := ""

	for _, s := range split {
		rVar := strings.ToLower(s)
		if _, ok := (*returnVars)[rVar]; ok {
			rVar += rVar[:1]
		}

		if unicode.IsUpper(rune(s[0])) {
			query += fmt.Sprintf("-[%s:%s]->", rVar, s)
		} else {
			query += fmt.Sprintf("(%s:%s)", rVar, strings.ToUpper(s[:1])+s[1:])
		}
	}

	if strings.HasSuffix(query, "->") {
		return query + "()"
	}
	return query
}
