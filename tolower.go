package piscine

func ToLower(s string) string {
	var result []rune
	for _, s := range s {
		if s >= 'A' && s <= 'Z' {
			result = append(result, s+32)
		} else {
			result = append(result, s)
		}
	}
	return string(result)
}
