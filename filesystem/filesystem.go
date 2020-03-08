package filesystem

import "os"

// FileExists checks the filesystem for the file.
// Returns true if the file exists; false otherwise.
func FileExists(f string) bool {
	if _, err := os.Stat(f); err == nil {
		return true
	}
	return false
}
