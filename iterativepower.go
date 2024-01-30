package piscine

func iterativePower(nb int, power int) int {
	result := nb
	if power == 0 {
		return 0
	}
	for i := 1; i < power+1; i++ {
		result *= nb
	}
	return result
}
