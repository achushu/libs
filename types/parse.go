package types

import (
	"errors"
	"strconv"
)

// ParseFloat64 converts an interface{} to string before parsing the float value
func ParseFloat64(v interface{}) (float64, error) {
	s := AssertString(v)
	if len(s) == 0 {
		return 0, errors.New("empty string")
	}
	return strconv.ParseFloat(s, 64)
}

func ParseInt(v interface{}) (int, error) {
	s := AssertString(v)
	if len(s) == 0 {
		return 0, errors.New("empty string")
	}
	i, err := strconv.ParseInt(s, 10, 32)
	return int(i), err
}

// ParseInt64 converts an interface{} to string before parsing the int value
func ParseInt64(v interface{}) (int64, error) {
	s := AssertString(v)
	if len(s) == 0 {
		return 0, errors.New("empty string")
	}
	return strconv.ParseInt(s, 10, 64)
}

func ParseBoolSlice(v interface{}) ([]bool, error) {
	s := AssertSlice(v)
	if len(s) == 0 {
		return nil, errors.New("no elements in slice")
	}
	b := make([]bool, len(s))
	for i, v := range s {
		b[i] = AssertBool(v)
	}
	return b, nil
}
