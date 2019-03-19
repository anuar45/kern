package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	lineIndex := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Println(os.Stderr, "dup2: %v\n", err)
			}
			countLines(f, counts)
			indexFile(f, lineIndex)
			f.Close()
		}

		for line, n := range counts {
			fmt.Printf("%d\t%s\t%s\n", n, line, lineIndex[line])
		}
	}
}

func countLines(f *os.File, count map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		count[input.Text()]++
	}
}

func indexFile(f *os.File, lineIndex map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if !strings.Contains(lineIndex[input.Text()], f.Name()) {
			lineIndex[input.Text()] += f.Name()

		}
	}
}
