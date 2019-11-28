package trim

import (
	"unicode"
)

func myTrim(b []byte) []byte {
	i, j := 0, 0

	for j < len(b) {
		if !unicode.IsSpace(rune(b[i])) || !unicode.IsSpace(rune(b[j])) {
			i++
			b[i] = b[j]
		}
		j++
	}
	return b[:i+1]
}
