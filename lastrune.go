package piscine

func LastRune(s string) rune {
	if len(s) > 0 {
		values := []rune(s)
		return values[len(s)-1]
	}
	return 0
}
