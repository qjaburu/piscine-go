package main

import "github.com/01-edu/z01"

func main() {
	for s := '0'; s <= '9'; s++ {
		z01.PrintRune(s)
	}
	z01.PrintRune('\n')
}
