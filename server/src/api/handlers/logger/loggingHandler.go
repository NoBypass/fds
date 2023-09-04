package logger

import (
	"fmt"
	"strings"
	"time"
)

const (
	ERROR   = "\033[31m"
	WARN    = "\033[33m"
	SUCCESS = "\033[32m"
	INFO    = "\033[34m"
	GREY    = "\033[90m"
	RESET   = "\033[0m"
)

func Log(msg string, color string) {
	currentTime := time.Now()
	formattedTime := currentTime.Format("15.04.05")
	lines := strings.Split(msg, "\n")
	for i := 0; i < len(lines); i++ {
		if i != 0 {
			lines[i] = fmt.Sprintf("           %s", lines[i])
		}
	}
	msg = strings.Join(lines, "\n")
	fmt.Printf("\u001B[1m(%s%s\033[37m) \u001B[0m%s", color, formattedTime, msg)
}
