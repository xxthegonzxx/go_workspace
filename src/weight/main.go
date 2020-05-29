// weight converts its numeric argument to Kilograms and Pounds.
package main

import (
	"fmt"
	"os"
	"strconv"
	weight "weight/conv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		p := weight.Pounds(t)
		k := weight.Kilograms(t)
		fmt.Printf("%s = %s, %s = %s\n",
			p, weight.PToK(p), k, weight.KToP(k))
	}
}

//!-
