package rev

import "testing"

func TestEmpty(t *testing.T) {
	want := []int{}
	got := []int{}
	Rev(got)
	if !cmpSlice(want, got) {
		t.Errorf("%v != %v\n", want, got)
	}
}

func TestRev(t *testing.T) {
	want := []int{1, 2, 4, 3}
	got := []int{3, 4, 2, 1}
	Rev(got)
	if !cmpSlice(want, got) {
		t.Errorf("%v != %v\n", want, got)
	}
}

func TestRevFixed(t *testing.T) {
	want := [...]int{1, 2, 4, 3, 7}
	got := [...]int{7, 3, 4, 2, 1}
	RevFixed(&got)
	if want != got {
		t.Errorf("%v != %v\n", want, got)
	}
}

func cmpSlice(a, b []int) bool {
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
