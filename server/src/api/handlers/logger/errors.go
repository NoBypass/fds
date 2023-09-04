package logger

import (
	"fmt"
	"runtime"
	"server/src/utils"
	"strings"
)

func Error(err error) {
	if err == nil {
		return
	}

	stackTrace := make([]byte, 4096)
	stackSize := runtime.Stack(stackTrace, false)
	stackTrace = stackTrace[:stackSize]

	errorMessage := fmt.Sprintf("%sError: %s%s\n", ERROR, RESET, err.Error())
	debugInfo := fmt.Sprintf("%sDebug Info:\n", ERROR)
	lines := strings.Split(string(stackTrace), "\n")
	for _, line := range lines {
		line = utils.TrimStringUntil(line, "fds/")
		if strings.HasPrefix(line, "fds/") {
			line = "  " + line
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			debugInfo += fmt.Sprintf("%s  %s%s%s:%s\n", GREY, RESET, parts[0], GREY, parts[1])
		}
	}

	Log(errorMessage+debugInfo, ERROR)
}
