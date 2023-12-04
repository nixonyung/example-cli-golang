package logger

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

type Logger struct {
	NoLoc     bool
	LocOffset int
}

var (
	Default = &Logger{}
)

func (l *Logger) output(logType logType, message string) {
	var codeLocation string // relative path starting with src
	if !(l.NoLoc) {
		// (ref.) [How do you get a Golang program to print the line number of the error it just called?](https://stackoverflow.com/a/24809646)
		if _, file, line, ok := runtime.Caller(2 + l.LocOffset); ok { // 2 for skipping this function (`output`) and the API (`Print`/`Fatal`)
			// `file` should be `/go/src/...` in the docker container, but we want to show as `src/...`
			fileTokens := strings.SplitN(file, "/", 3)
			codeLocation = fmt.Sprintf("%s:%d", fileTokens[2], line)
		} else {
			codeLocation = "???"
		}
	}

	fmt.Fprintf(os.Stderr, "%s (%s) %s\n",
		logType,
		codeLocation,
		message,
	)
}

func (l *Logger) Printf(format string, a ...any) {
	l.output(logTypeInfo, fmt.Sprintf(format, a...))
}

func (l *Logger) Fatal(err error) {
	l.output(logTypeError, err.Error())
	os.Exit(1)
}
