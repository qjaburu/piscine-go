package piscine

func FirstRune(s string) rune {
	if len(s) > 0 {
		letters := []rune(s)
		return letters[0]
	}
	return 0
}
