package comma

import "testing"

func TestComma(t *testing.T) {
	want := "12,345"
	got := comma("12345")
	if want != got {
		t.Errorf("want %s but got %s\n", want, got)
	}
}

func TestShort(t *testing.T) {
	want := "123"
	got := comma("123")
	if want != got {
		t.Errorf("want %s but got %s\n", want, got)
	}
}
