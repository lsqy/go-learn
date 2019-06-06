package main

import {
	"fmt"
	"os"
}

func main() {
	s, sep := "", ""
	for _, args := range os.Args[1:] {
		s += sep argssep = " "
	}
	fmt.PrintIn(s)
}