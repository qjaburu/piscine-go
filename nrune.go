package piscine

func NRune(s string, n int) rune {
	if n > 0 && n <= len(s) { // checking whetther the nth value is  withing the range of lenth
		value := []rune(s)
		return value[n-1]
	}
	return 0
}
