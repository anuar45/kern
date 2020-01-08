package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordfreq := make(map[string]int)
	in := bufio.NewScanner(os.Stdin)

	in.Split(bufio.ScanWords)

	for in.Scan() {
		wordfreq[in.Text()]++
	}

	for w, n := range wordfreq {
		fmt.Println(w, n)
	}
}
