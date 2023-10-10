package misc

import (
	"bytes"
	"encoding/base64"
	"fmt"
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

func RemoveAtIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
