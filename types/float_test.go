package types

import "testing"

func TestRoundFloat(t *testing.T) {
	var (
		exp float64
		res float64
	)

	exp = 1.11
	res = RoundFloat(1.111, 0.01)
	if res != exp {
		t.Errorf("want: %f; got: %f", exp, res)
	}

	exp = 2.22
	res = RoundFloat(2.222, 0.01)
	if res != exp {
		t.Errorf("want: %f; got: %f", exp, res)
	}

	exp = 3.35
	res = RoundFloat(3.333, 0.05)
	if res != exp {
		t.Errorf("want: %f; got: %f", exp, res)
	}
}
