package main 

import "github.com/01-edu/z01"
 
func PrintComb() {

	for x := '0' ; x <= '7' ; x++ {

		for y := x+1 ; y<= '8' ; y++ {

			for z :=y+1 ; y<= '9' ; z++ {

			z01.PrintComb(x)
		   z01.PrintComb(y)
		   z01.PrintComb(z)


		   if x != '7' || y != '8' || z != '9'{
			z01.PrintRune(',')
			z01.PrintRune('')
			}
		}
	}
	z01.PrintRune('\n')
}