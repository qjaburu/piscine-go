package main

import "fmt"

func main() {
	for i := 0; i <= 7; i++ {
		for j := i + 1; j <= 8; j++ {
			for k := j + 1; k <= 9; k++ {

				fmt.Print(i)
				fmt.Print(j)
				fmt.Print(k)

				if i != 7 || j != 8 || k != 9 {
					fmt.Print(", ")
				}

			}
		}
	}
	fmt.Println()
}
