package anagram

import "testing"

func TestAnagramEqual(t *testing.T) {
	equal := checkAnagram("aabbccc", "cccabba")
	if !equal {
		t.Error("failed to check equal anagrams, should return true")
	}
}

func TestAnagramEmpty(t *testing.T) {
	equal := checkAnagram("", "")
	if !equal {
		t.Error("failed to check empty anagrams, should return true")
	}
}

func TestAnagramNonEqual(t *testing.T) {
	equal := checkAnagram("abc", "aba")
	if equal {
		t.Error("failed to check non equal anagrams, should return false")
	}
}
