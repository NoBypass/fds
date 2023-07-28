package dbutils

import (
	"encoding/base64"
	"fmt"
	"time"
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

func GenerateJoinedAt() int64 {
	return time.Now().UnixMilli()
}
