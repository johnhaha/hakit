package haerr

import (
	"fmt"
	"runtime"
)

func GetTrace() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return fmt.Sprintf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
}

func Trace() {
	fmt.Printf(GetTrace())
}
