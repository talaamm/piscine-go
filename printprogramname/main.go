package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	A := os.Args[0][2:]
	for _, i := range A {
		z01.PrintRune(i)
	}

	z01.PrintRune('\n')
}
