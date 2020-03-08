package types

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestAtoi(t *testing.T) {
	expected := rand.Intn(10000)
	res := Atoi(strconv.Itoa(expected))
	if expected != res {
		t.Errorf("unexpected result -- want: %v; got: %v\n", expected, res)
	}

	expected = -1
	res = Atoi("NaN")
	if expected != res {
		t.Errorf("unexpected result -- want: %v; got: %v\n", expected, res)
	}
}

func TestBtoMB(t *testing.T) {
	expected := int64(1)
	res := BtoMB(1048576)
	if expected != res {
		t.Errorf("unexpected result -- want: %v; got: %v\n", expected, res)
	}
}
