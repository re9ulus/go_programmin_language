package freq

import "testing"

func TestFreq(t *testing.T) {
	want := map[string]int{
		"yolo": 2,
		"this": 1,
		"is":   1,
		"test": 3,
	}

	got := freq("yolo this is test test yolo test")

	if !cmpMap(want, got) {
		t.Errorf("%v != %v\n", want, got)
	}
}

func cmpMap(want, got map[string]int) bool {
	if len(want) != len(got) {
		return false
	}

	for key, val1 := range want {
		if val2, ok := got[key]; !ok || val1 != val2 {
			return false
		}
	}
	return true
}
