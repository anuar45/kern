package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/anuar45/kern/ch2/popcount"
)

func main() {
	a, _ := strconv.Atoi(os.Args[1])
	fmt.Println(popcount.PopCount(uint64(a)))

}
