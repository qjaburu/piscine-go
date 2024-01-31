package piscine

func NRune(s string, n int) rune {
	if len(s) > 0 {
		numbers := []rune(s)
		return numbers[n]
	}
	return 0
}
