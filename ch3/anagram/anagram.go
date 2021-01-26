package anagram

// Task 3.12 check if two strings are anagrams
func checkAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	d := make(map[rune]int)

	for _, c := range s1 {
		d[c]++
	}
	for _, c := range s2 {
		d[c]--
	}

	for _, v := range d {
		if v != 0 {
			return false
		}
	}

	return true
}
