package types

import "strconv"

// Atoi converts a string to an int. Returns -1 if the string is not a number.
func Atoi(a string) int {
	if i, err := strconv.Atoi(a); err == nil {
		return i
	}
	return -1
}

// BtoMB converts bytes to megabytes
func BtoMB(n int64) int64 {
	return n / 1024 / 1024
}
