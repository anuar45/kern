package main

import (
	"fmt"
	"io"
	"os"
	"bufio"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	var utfLen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for r, n, err := in.ReadRune() {
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, )
		}
	}

	fmt.Println("vim-go")
}
