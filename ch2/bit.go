package main

import "fmt"

func main() {
	var x uint8 = 5

	fmt.Println(x >> 1)
	fmt.Println(x & 1)
	fmt.Println(x&1 == 1)
}
