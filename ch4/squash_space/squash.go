package squash

import "unicode"

func squash(s []rune) []rune {
	if len(s) < 2 {
		return s
	}
	i := 1
	isLastSpace := unicode.IsSpace(s[0])
	for _, v := range s[1:] {
		isSpace := unicode.IsSpace(v)
		if isSpace && isLastSpace {
			continue
		}
		s[i] = v
		isLastSpace = isSpace
		i++
	}

	return s[:i]
}
