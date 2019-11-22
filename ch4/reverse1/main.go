package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5}

	reverse(sl)

	fmt.Println(sl)
}

func reverse(sl []int) {
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
}
