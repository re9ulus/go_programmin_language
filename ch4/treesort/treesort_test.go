package treesort

import "testing"

func TestSort(t *testing.T) {
	want := []int{1, 2, 3, 4, 5}
	got := []int{5, 3, 4, 1, 2}
	Sort(got)
	if !cmpSlice(want, got) {
		t.Errorf("%v != %v\n", want, got)
	}
}

func cmpSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
