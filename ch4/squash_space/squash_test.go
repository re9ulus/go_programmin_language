package squash

import "testing"

func TestSquash(t *testing.T) {
	want := []rune("hello world! check")
	got := []rune("hello   world!  check")
	got = squash(got)
	if !cmpSlice(got, want) {
		t.Errorf("%s != %s", string(want), string(got))
	}
}

func cmpSlice(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
