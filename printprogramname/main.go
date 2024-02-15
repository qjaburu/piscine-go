package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	if len(arg) > 0 {
		arg := os.Args[0][2:]
		argument := []rune(arg)

		for _, c := range argument {
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')
}
