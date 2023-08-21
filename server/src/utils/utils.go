package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
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

func convertCamelToSnake(input string) string {
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

func JoinMap(input map[string]string, separator string) string {
	var contents []string
	for _, value := range input {
		contents = append(contents, value)
	}

	return strings.Join(contents, separator)
}

func InsertAtIndex(arr []string, index int, element string) []string {
	newArr := make([]string, len(arr)+1)

	copy(newArr[:index], arr[:index])
	newArr[index] = element
	copy(newArr[index+1:], arr[index:])

	return newArr
}

func GetMapIndex[T interface{}](input map[string]T, key string) int {
	current := 0
	for k, _ := range input {
		if k == key {
			return current
		}
		current++
	}

	return -1
}
