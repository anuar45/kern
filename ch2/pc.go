package main

import (
	"fmt"
	"github.com/anuar45/kern/ch2/popcount"
	"os"
	"strconv"
	"time"
)

func main() {
	a, _ := strconv.Atoi(os.Args[1])
	var start time.Time

	start = time.Now()
	fmt.Println(popcount.PopCount(uint64(a)))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println(popcount.PopCountFor(uint64(a)))
	fmt.Println(time.Since(start))

}
