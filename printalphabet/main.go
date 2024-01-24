package main

import "github.com/01-edu/z01"

func main() {
	for q := 'a'; q <= 'z'; q++ {
		z01.PrintRune(q)
	}
	z01.PrintRune('\n')
}
