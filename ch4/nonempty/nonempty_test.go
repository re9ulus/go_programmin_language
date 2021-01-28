package nonempty

import "testing"

func TestNonEmpty(t *testing.T) {
	want := []string{"hello", "world"}
	got := nonempty([]string{"", "hello", "", "", "world", ""})
	if !cmpSlices(want, got) {
		t.Errorf("%v != %v\n", want, got)
	}
}

func TestNonEmpty2(t *testing.T) {
	want := []string{"hello", "world"}
	got := nonempty2([]string{"", "hello", "", "", "world", ""})
	if !cmpSlices(want, got) {
		t.Errorf("%v != %v\n", want, got)
	}
}

func cmpSlices(a, b []string) bool {
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
