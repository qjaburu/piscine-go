package piscine

func NRune(s string, n int) rune {
	if len(s) > 0 {
		r := []rune(s)
		return r[n]
	}
	return 0
}
