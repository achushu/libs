package types

import (
	"math/rand"
	"reflect"
	"strconv"
	"testing"
)

func TestAssertBoolSlice(t *testing.T) {
	expected := []bool{true, false, false, true}
	res := AssertBoolSlice(expected)
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("unexpected result -- want: %v; got: %v\n", expected, res)
	}

	expected = []bool{}
	res = AssertBoolSlice("test")
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("unexpected result -- want: %v; got: %v\n", expected, res)
	}
}

func TestAssertByteSlice(t *testing.T) {
	expected := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := AssertByteSlice(expected)
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("unexpected result -- want: %v; got: %v\n", expected, res)
	}

	expected = []byte{}
	res = AssertByteSlice("test")
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("unexpected result -- want: %v; got: %v\n", expected, res)
	}
}

func TestAssertInt(t *testing.T) {
	expected := rand.Intn(10000)
	res := AssertInt(expected)
	if expected != res {
		t.Errorf("unexpected result -- want: %v; got: %v\n", expected, res)
	}

	expected = 0
	res = AssertInt("NaN")
	if expected != res {
		t.Errorf("unexpected result -- want: %v; got: %v\n", expected, res)
	}
}

func TestAssertFloat64(t *testing.T) {
	expected := rand.Float64() * 10.0
	res := AssertFloat64(expected)
	if expected != res {
		t.Errorf("unexpected result -- want: %v; got: %v\n", expected, res)
	}

	expected = 0
	res = AssertFloat64("NaN")
	if expected != res {
		t.Errorf("unexpected result -- want: %v; got: %v\n", expected, res)
	}
}

func TestAssertString(t *testing.T) {
	expected := strconv.FormatInt(rand.Int63n(10000), 16)
	res := AssertString(expected)
	if expected != res {
		t.Errorf("unexpected result -- want: %v; got: %v\n", expected, res)
	}

	expected = ""
	res = AssertString(rand.Intn(10000))
	if expected != res {
		t.Errorf("unexpected result -- want: %v; got: %v\n", expected, res)
	}
}

func TestAssertStringSlice(t *testing.T) {
	expected := []string{"a", "b", "c", "d", "e", "f"}
	res := AssertStringSlice(expected)
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("unexpected result -- want: %v; got: %v\n", expected, res)
	}

	expected = []string{}
	res = AssertStringSlice("test")
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("unexpected result -- want: %v; got: %v\n", expected, res)
	}
}
