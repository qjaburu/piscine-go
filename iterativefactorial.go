package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0 // returns when n is negative
	}
	result := 1
	for i := 1; i <= nb; i++ {
		if result*i < result {
			return 0
		}
		result *= i
	}
	return result
}
