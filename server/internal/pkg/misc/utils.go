package misc

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"reflect"
	"strings"
	"time"
	"unicode"
)

func GenerateUUID(args ...any) string {
	currentTime := time.Now().UnixNano()
	bytes := []byte{byte(currentTime)}

	for _, arg := range args {
		bytes = append(bytes, []byte(fmt.Sprintf("%v", arg))...)
	}

	encoded := base64.StdEncoding.EncodeToString(bytes)
	return encoded
}

func GetNowInMs() int64 {
	return time.Now().UnixMilli()
}

func ConvertCamelToSnake(input string) string {
	var output bytes.Buffer

	for i, char := range input {
		if unicode.IsUpper(char) {
			if i > 0 {
				output.WriteRune('_')
			}
			output.WriteRune(unicode.ToLower(char))
		} else {
			output.WriteRune(char)
		}
	}

	return output.String()
}

func FirstLower(input string) string {
	return strings.ToLower(input[:1]) + input[1:]
}

func FirstUpper(input string) string {
	return strings.ToUpper(input[:1]) + input[1:]
}

func MaxOutAt(input int64, max int64) int64 {
	if input > max {
		return max
	}
	return input
}

func TrimStringUntil(s, substr string) string {
	index := strings.Index(s, substr)
	if index == -1 {
		return s
	}
	return s[index:]
}

func MapResult[T any](input *T, result *neo4j.EagerResult, indexLetter string) (*T, error) {
	if len(result.Records) == 0 {
		return nil, fmt.Errorf("no results found")
	}
	r, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], indexLetter)
	if err != nil {
		return nil, err
	}

	inputValue := *input
	inputType := reflect.TypeOf(inputValue)
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		fieldType := field.Type
		jsonTag := field.Tag.Get("json")

		var val interface{}
		switch fieldType.Kind() {
		case reflect.String:
			val, err = neo4j.GetProperty[string](r, jsonTag)
		case reflect.Int64:
			val, err = neo4j.GetProperty[int64](r, jsonTag)
		case reflect.Bool:
			val, err = neo4j.GetProperty[bool](r, jsonTag)
		case reflect.Float64:
			val, err = neo4j.GetProperty[float64](r, jsonTag)
		default:
			continue
		}
		if err != nil {
			return nil, err
		}

		reflect.ValueOf(input).Elem().Field(i).Set(reflect.ValueOf(val))
	}
	*input = inputValue
	return input, nil
}

func StructToMap[T any](input *T) map[string]any {
	values := make(map[string]any)
	inputType := reflect.TypeOf(*input)
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := reflect.ValueOf(input).Elem().FieldByName(field.Name).Interface()
		if reflect.DeepEqual(value, reflect.Zero(field.Type).Interface()) {
			continue
		}
		values[ConvertCamelToSnake(field.Name)] = value
	}
	return values
}
