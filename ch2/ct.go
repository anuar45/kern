// Cf converts its numeric arguments to Celsius and Fahrenheit
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/anuar45/kern/ch2/multiconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f := multiconv.Fahrenheit(t)
		c := multiconv.Celsius(t)
		kg := multiconv.Kilo(t)
		p := multiconv.Pound(t)
		m := multiconv.Meter(t)
		ft := multiconv.Feet(t)

		fmt.Printf("%s = %s, %s = %s\n", f, multiconv.FToC(f), c, multiconv.CToF(c))

		fmt.Printf("%s = %s, %s = %s\n", kg, multiconv.KToP(kg), p, multiconv.PToK(p))

		fmt.Printf("%s = %s, %s = %s\n", m, multiconv.MToF(m), ft, multiconv.FToM(ft))

	}
}
