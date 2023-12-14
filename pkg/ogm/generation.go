package ogm

import (
	"fmt"
	"strings"
	"unicode"
)

func generateCypherQuery(preloads []string) (map[string]string, map[string][]string, map[string]string) {
	queries := make(map[string]string)
	returnVars := make(map[string][]string)
	varNames := make(map[string]string)

	for _, preload := range preloads {
		split := strings.Split(preload, ".")

		varName := strings.ToLower(split[len(split)-2][:1]) + strings.Replace(fmt.Sprintf("%d", len(split)-1), "1", "", 1)
		varVal := strings.ToLower(split[len(split)-1])
		returnVars[varName] = append(returnVars[varName], varVal)
		varNames[split[len(split)-2]] = varName

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
		queries[k] = generator(k, varNames)
	}
	return queries, returnVars, varNames
}

func generator(preload string, varNames map[string]string) string {
	split := strings.Split(preload, ".")
	query := ""

	for _, s := range split {
		name := varNames[s]

		if unicode.IsUpper(rune(s[0])) {
			query += fmt.Sprintf("-[%s:%s]->", name, s)
		} else {
			query += fmt.Sprintf("(%s:%s)", name, strings.ToUpper(s[:1])+s[1:])
		}
	}

	if strings.HasSuffix(query, "->") {
		return query + "()"
	}
	return query
}
