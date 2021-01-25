package basename

import "testing"

func TestBasename(t *testing.T) {
	want := "name"
	got := basename("/simple/name.txt")
	if want != got {
		t.Errorf("want %s but got %s", want, got)
	}
}

func TestSlash(t *testing.T) {
	want := ""
	got := basename("/simple/name/")
	if want != got {
		t.Errorf("want %s but got %s", want, got)
	}
}

func TestEmpty(t *testing.T) {
	want := ""
	got := basename("")
	if want != got {
		t.Errorf("want %s but got %s", want, got)
	}
}
