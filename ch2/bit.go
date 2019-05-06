package main

import "fmt"

func main() {
	var x uint8 = 7

	fmt.Printf("%b\n", x)
	fmt.Println(x >> 1)
	fmt.Println(x & 1)
	fmt.Println(x&1 == 1)
	fmt.Println(x & (x - 1))
}
