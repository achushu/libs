package filesystem

import "testing"

func TestFileExists(t *testing.T) {
	this := "filesystem.go"
	if !FileExists(this) {
		t.Errorf("expected %s to exist", this)
	}

	fake := "abc.xyz.123"
	if FileExists(fake) {
		t.Errorf("expected %s to not exist", fake)
	}
}
