// Initialize a globally available "console" to log to.
// Writes out to stdout and by default only logs messages
// INFO or higher (can be configured).

package out

import (
	"fmt"
	"log"
)

var (
	stdCfg = &Config{
		Enabled:   true,
		Filename:  "",
		Async:     false,
		Threshold: PriorityDebug,
		Rotate:    nil,
	}
	stdout *Log
)

func init() {
	var err error
	stdout, err = New(stdCfg)
	if err != nil {
		log.Fatalf("error initializing global output: %s", err)
	}
}

// Errorln is a globally available logging function
// at the INFO level.
func Errorln(args ...interface{}) {
	stdout.Logln(PriorityError, fmt.Sprint(args...))
}

// Errorf is a globally available logging function
// at the INFO level.
func Errorf(format string, args ...interface{}) {
	stdout.Logf(PriorityError, format, args...)
}

// Println is a globally available logging function
// at the INFO level.
func Println(args ...interface{}) {
	stdout.Logln(PriorityInfo, fmt.Sprint(args...))
}

// Printf is a globally available logging function
// at the INFO level.
func Printf(format string, args ...interface{}) {
	stdout.Logf(PriorityInfo, format, args...)
}

// Write is a globally available output function
func Write(msg []byte) {
	stdout.Write(msg)
}

// WriteString is a globally available output function
func WriteString(msg string) {
	stdout.WriteString(msg)
}
