package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	for index, value := range os.Args {
		fmt.Println(index, value)
	}
	fmt.Println(time.Since(start).Seconds())
}
