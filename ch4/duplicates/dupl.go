package dupl

func dedupl(ar []int) []int {
	if len(ar) < 2 {
		return ar
	}

	i := 1
	last := ar[0]
	for _, v := range ar {
		if v != last {
			ar[i], last = v, v
			i++
		}
	}
	return ar[:i]
}
