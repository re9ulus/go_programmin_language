package tempconv

import (
	"math"
	"testing"
)

const eps = 1e-6

func isAlmostEqual(a, b float64) bool {
	return math.Abs(a-b) < eps
}

func TestCtoF(t *testing.T) {
	c := Celsius(14.0)
	got := CtoF(c)
	want := Fahrenheit(57.2)
	if !isAlmostEqual(float64(got), float64(want)) {
		t.Errorf("%g != %g\n", got, want)
	}
}

func TestFtoC(t *testing.T) {
	f := Fahrenheit(76)
	got := FtoC(f)
	want := Celsius(24.444444444444443)
	if !isAlmostEqual(float64(got), float64(want)) {
		t.Errorf("%g != %g\n", got, want)
	}
}
