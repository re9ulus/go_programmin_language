package appendImpl

import "testing"

func TestAppend(t *testing.T) {
	want := []int{1, 2, 3, 4}
	got := appendInt([]int{1, 2, 3}, 4)
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
