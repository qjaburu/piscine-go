package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	set := false
	for x := '0'; x <= '9'; x++ {
		for y := x + 1; y <= '9'; y++ {
			for z := y + 1; z <= '9'; z++ {

				if set {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
				set = true
				z01.PrintRune(x)
				z01.PrintRune(y)
				z01.PrintRune(z)
			}
		}
	}

	z01.PrintRune('\n')
}
