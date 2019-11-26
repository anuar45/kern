package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
}

func strDup(sl []string) []string {
	i, j := 0, 0
	for j < len(sl) {
		if sl[i] != sl[j] {
			sl[i] = sl[j]
			i++
			j++
		} else {
			j++
		}
	}

	return sl[:i]
}
