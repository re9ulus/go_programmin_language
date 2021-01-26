package comma

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// task 3.10
// build nonrecursive comma with butest.Buffer
// task 3.11
// enhance comma to support floating-point and optional sign
func comma2(s string) string {
	const n = 3
	var sign string
	var floatPart string
	if len(s) > 0 && (s[0] == '+' || s[0] == '-') {
		sign = string(s[0])
		s = s[1:]
	}
	if dot := strings.Index(s, "."); dot >= 0 {
		floatPart = s[dot:]
		s = s[:dot]
	}

	var buf bytes.Buffer
	c := 0
	for i := len(s) - 1; i >= 0; i, c = i-1, c+1 {
		if c == n {
			buf.WriteByte(',')
			c = 0
		}
		fmt.Fprintf(&buf, "%c", s[i])
	}
	if len(sign) == 1 {
		buf.WriteString(sign)
	}
	ans := buf.Bytes()
	reverse(ans)
	return string(ans) + floatPart
}

func reverse(buf []byte) {
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
}
