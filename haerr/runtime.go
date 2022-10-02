package haerr

import (
	"fmt"
	"runtime"
)

/* trace skip
0: caller it self, no use

1: GetTrace func, no use

2: place call GetTrace

3: func call GetTrace
*/
func GetTrace(skip int) string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(skip, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return fmt.Sprintf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
}

func Trace() {
	fmt.Printf(GetTrace(3))
}
