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

func MaxOutAt(input int64, max int64) int64 {
	if input > max {
		return max
	}

	return input
}
