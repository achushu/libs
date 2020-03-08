// +build !debug

package out

// Debugln is a no-op when not compiled with a debug tag.
func Debugln(args ...interface{}) {
	// do nothing
}

// Debugf is a no-op when not compiled with a debug tag.
func Debugf(format string, args ...interface{}) {
	// do nothing
}
