package dupl

import "testing"

// Task 4.5 inplace eliminate adjacent duplicates
func TestDupl(t *testing.T) {
	want := []int{1, 2, 4}
	got := []int{1, 1, 1, 2, 2, 4}
	got = dedupl(got)
	if !cmpSlice(want, got) {
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
