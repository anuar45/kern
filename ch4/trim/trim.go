package trim

import (
	"encoding/utf8"
	"fmt"
)

func myTrim(b []byte) []byte {
	i, j := 0, 0

	for j < len(b) {
		if utf8.IsSpace(b[i]) {

			if utf8.IsSpace(b[j]) {
				j++
				continue
			}
		}
		b[i] = b[j]
		i++
	}
}
