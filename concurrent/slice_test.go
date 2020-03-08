package concurrent

import (
	"reflect"
	"strconv"
	"testing"
)

func TestStringSlice(t *testing.T) {
	cap := 10
	s := NewStringSlice(cap)
	if s.Cap() != cap {
		t.Errorf("incorrect capacity -- want: %d; got: %d\n", cap, s.Cap())
		t.FailNow()
	}
	s.Add("test")
	if s.Len() != 1 {
		t.Errorf("incorrect length -- want: %d; got: %d\n", 1, s.Len())
	}
	res := s.Read(false)
	s.ReadDone()
	if len(res) != 1 {
		t.Errorf("incorrect number of items -- want: %d; got: %d\n", cap, len(res))
	}
}

func TestStringSliceFill(t *testing.T) {
	cap := 10
	expected := make([]string, 0, cap)
	for i := 0; i < cap; i++ {
		expected = append(expected, strconv.Itoa(i))
	}
	s := NewStringSlice(cap)
	for _, v := range expected {
		s.Add(v)
	}
	if s.Len() != s.Cap() {
		t.Errorf("incorrect length -- want: %d; got: %d\n", cap, s.Len())
	}
	res := s.Read(false)
	s.ReadDone()
	if len(res) != cap {
		t.Errorf("incorrect number of items -- want: %d; got: %d\n", cap, len(res))
	}
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("result did not match expected slice:\nwant: %v\ngot: %v\n", expected, res)
	}
}

func TestStringSliceOverfill(t *testing.T) {
	cap := 10
	s := NewStringSlice(cap)
	for i := 0; i < cap; i++ {
		s.Add(strconv.Itoa(i))
	}
	res := s.Add("one more")
	if res {
		t.Error("overcapacity add did not fail as expected")
	}
}

func TestStringSliceClear(t *testing.T) {
	cap := 10
	expected := make([]string, 0, cap)
	for i := 0; i < cap; i++ {
		expected = append(expected, strconv.Itoa(i))
	}
	s := NewStringSlice(cap)
	for _, v := range expected {
		s.Add(v)
	}
	if s.Len() != s.Cap() {
		t.Errorf("incorrect length -- want: %d; got: %d\n", cap, s.Len())
	}
	res := s.Read(true)
	s.ReadDone()
	if len(res) != cap {
		t.Errorf("incorrect number of items -- want: %d; got: %d\n", cap, len(res))
	}
	if s.Len() != 0 {
		t.Error("slice did not clear")
	}
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("result did not match expected slice:\nwant: %v\ngot: %v\n", expected, res)
	}
	res = s.Read(false)
	s.ReadDone()
	if len(res) != 0 {
		t.Errorf("read returned a non-empty result after clear\n%v\n", res)
	}
	s.Clear()
}
