package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
}

func strDup(sl []string) []string {
	i, j := 0, 1
	for j < len(sl) {
		if sl[i] != sl[j] {
			i++
			sl[i] = sl[j]
		}
		j++
	}

	return sl[:i+1]
}
