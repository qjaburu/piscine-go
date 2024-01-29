package main

import "fmt"

func main() {
	QuadA(5, 1)
}

func QuadA(x, y int) {
	if !((x <= 0 || y <= 0) || (x <= 0 && y <= 0)) {
		PrintFirstLine(x)

		for i := 0; i < (y - 2); i++ {
			PrintSecondLine(x)
		}
		PrintLastLine(x)

	}
}

func PrintFirstLine(x int) {
	for i := 0; i < x; i++ {
		if i == 0 || i == (x-1) {
			fmt.Printf("%c", 'o')
		} else {
			fmt.Printf("%c", '-')
		}
	}
	fmt.Println()
}

func PrintSecondLine(x int) {
	for i := 0; i < x; i++ {
		if i == 0 || i == x-1 {
			fmt.Printf("%c", '|')
		} else {
			fmt.Printf("%c", ' ')
		}
	}
	fmt.Println()
}

func PrintLastLine(x int) {
	for i := 0; i < x; i++ {
		if i == 0 || i == (x-1) {
			fmt.Printf("%c", 'o')
		} else {
			fmt.Printf("%c", '-')
		}
	}
	fmt.Println()
}
