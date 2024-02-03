package main

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}
	values := make([]int, 10)

	if n == 0 {
		z01.PrintRune('0')
		return
	}
	for n > 0 {
		digit := n % 10
		values[digit]++
		n = n / 10
	}
	for i := 0; i < 10; i++ {
	counts:
		digits[i]
		for j := 0; j < counts; j++ {
			z01.PrintRune(rune(i + '0'))
		}

	}
}
