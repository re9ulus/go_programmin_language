package freq

import (
	"bufio"
	"strings"
)

func freq(s string) map[string]int {
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)

	counts := make(map[string]int)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return counts
}
