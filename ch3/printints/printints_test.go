package printints

import "testing"

func TestPrintints(t *testing.T) {
	want := "[1, 2, 3]"
	got := printInts([]int{1, 2, 3})
	if want != got {
		t.Errorf("`%v` != `%v`\n", want, got)
	}
}
