package piscine

import "github.com/01-edu/z01"

func JumpOver(str string) string {
	var value string
	for i, char := range str {
		if (i+1)%3 == 0 {
			value += string(char)
		}
		if value == "" {
			z01.PrintRune('\n')
		}
	}
	return value + "\n"
}
