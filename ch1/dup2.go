package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	counts := make(map[string]int)
	lineIndex := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, lineIndex)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Println(os.Stderr, "dup2: %v\n", err)
			}
			countLines(f, counts, lineIndex)
			f.Close()
		}

		for line, n := range counts {
			fmt.Printf("%d\t%s\t%s\n", n, line, lineIndex[line])
		}
	}
}

func countLines(f *os.File, count map[string]int, lineIndex map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		count[input.Text()]++
		if !strings.Contains(lineIndex[input.Text()], f.Name()) {
			lineIndex[input.Text()] += filepath.Base(f.Name())
		}
	}
}
