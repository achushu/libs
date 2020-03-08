package types

import (
	"reflect"
	"testing"
)

func TestParseBoolSlice(t *testing.T) {
	expected := []bool{true, true, false, false}
	input := []string{"true", "true", "false", "false"}
	res, err := ParseBoolSlice(input)
	if err != nil {
		t.Errorf("error parsing slice: %s", err)
		t.FailNow()
	}
	if !reflect.DeepEqual(expected, res) {
		t.Errorf(" want: %v; got: %v\n", expected, res)
	}
}
