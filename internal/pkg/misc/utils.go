package misc

import (
	"strings"
	"time"
)

func JoinStrMap(m map[string]string, sep string) string {
	var str string
	for _, v := range m {
		str += v + sep
	}
	return str
}

func JoinStrArrMap(m map[string][]string, sep string) string {
	var str string
	for k, v := range m {
		for _, item := range v {
			str += k + "." + item + sep
		}
	}
	return str
}

func GetNowInMs() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func DeSnake(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, "_", ""))
}
