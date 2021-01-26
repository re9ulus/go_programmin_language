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

func TestNonRecComma(t *testing.T) {
	want := "12,345"
	got := comma2("12345")
	if want != got {
		t.Errorf("want %s but got %s\n", want, got)
	}
}

func TestNonRecShort(t *testing.T) {
	want := "123"
	got := comma2("123")
	if want != got {
		t.Errorf("want %s but got %s\n", want, got)
	}
}

func TestFloat(t *testing.T) {
	want := "-12,345.00023"
	got := comma2("-12345.00023")
	if want != got {
		t.Errorf("want %s but got %s\n", want, got)
	}
}
