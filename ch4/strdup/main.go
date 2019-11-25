package main

import (
		"fmt"
)

func main() {
		fmt.Println("Hello")

}

func strDup(sl []string) {
		for i,j := 0, 1; j < len(sl); i++ {
				if sl[i] == sl[j)] {
						sl[i] = sl[j]  
				}		
		}
}
