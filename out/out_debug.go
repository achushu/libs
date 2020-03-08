// +build debug

package out

// Functions are available when built with `-tags debug`

import "fmt"

// Debugln is a globally available logging function
// at the INFO level.
func Debugln(args ...interface{}) {
	stdout.Logln(PriorityDebug, fmt.Sprint(args...))
}

// Debugf is a globally available logging function
// at the INFO level.
func Debugf(format string, args ...interface{}) {
	stdout.Logf(PriorityDebug, format, args...)
}
