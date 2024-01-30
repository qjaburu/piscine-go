package piscine

import "fmt"

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}

	result := 1
	for i := 0; i < power; i++ {
		//result *= nb
		result = result * nb
	}
	return result
}
