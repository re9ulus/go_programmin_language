package intset

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var x IntSet
	x.Add(1)

	if !x.Has(1) {
		t.Errorf("insert failed")
	}
	if x.Has(2) {
		t.Errorf("testset lies")
	}
}

func TestUnion(t *testing.T) {
	var x, y IntSet
	x.Add(4)
	y.Add(9)

	x.UnionWith(&y)
	if !x.Has(4) {
		t.Errorf("union failed")
	}
	if !x.Has(9) {
		t.Errorf("union failed")
	}
}

func TestLen(t *testing.T) {
	var x IntSet
	if x.Len() != 0 {
		t.Errorf("wrong len, expect 0")
	}
	x.Add(5)
	if x.Len() != 1 {
		t.Errorf("wrong len, expect 1")
	}
	x.Add(5)
	if x.Len() != 1 {
		t.Errorf("wrong len, expect 1")
	}
	x.Add(70)
	if x.Len() != 2 {
		t.Errorf("wrong len, expect 2")
	}
}

func TestRemove(t *testing.T) {
	var x IntSet
	x.Add(42)
	if !x.Has(42) {
		t.Errorf("insert failed")
	}
	x.Remove(42)
	if x.Has(42) {
		t.Errorf("remove failed")
	}
}

func TestClean(t *testing.T) {
	var x IntSet
	x.Add(42)
	x.Clear()
	if x.Len() != 0 {
		t.Errorf("clear failed")
	}
}

func TestCopy(t *testing.T) {
	var x IntSet
	x.Add(42)
	y := x.Copy()
	if !y.Has(42) {
		t.Errorf("copy failed")
	}
	x.Add(5)
	if y.Has(5) {
		t.Errorf("copy failed, copy inset referenced same underlying array")
	}
}

func TestAddAll(t *testing.T) {
	var x IntSet
	x.AddAll(1, 2, 3, 4, 5)
	if !x.Has(4) {
		t.Errorf("add all failed")
	}
}
