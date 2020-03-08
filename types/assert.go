package types

// AssertBoolSlice ensures that the value is a bool slice
func AssertBoolSlice(v interface{}) []bool {
	if a, ok := v.([]bool); ok {
		return a
	}
	return []bool{}
}

// AssertByteSlice ensures that the value is a byte slice
func AssertByteSlice(v interface{}) []byte {
	if a, ok := v.([]byte); ok {
		return a
	}
	return []byte{}
}

func AssertBool(v interface{}) bool {
	if a, ok := v.(bool); ok {
		return a
	}
	return false
}

func AssertFloat64(v interface{}) float64 {
	switch a := v.(type) {
	case float32:
		return float64(a)
	case float64:
		return a
	default:
		return float64(AssertInt(v))
	}
}

// AssertInt ensures that the value is an int
func AssertInt(v interface{}) int {
	switch a := v.(type) {
	case int:
		return a
	case int32:
		return int(a)
	case int64:
		return int(a)
	case float64:
		return int(a)
	}
	return 0
}

func AssertSlice(v interface{}) []interface{} {
	if a, ok := v.([]interface{}); ok {
		return a
	}
	return []interface{}{}
}

// AssertString ensures that the value is a string
func AssertString(v interface{}) string {
	if a, ok := v.(string); ok {
		return a
	}
	return ""
}

func AssertStringSlice(v interface{}) []string {
	if a, ok := v.([]string); ok {
		return a
	}
	return []string{}
}
