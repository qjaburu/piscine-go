package piscine

func ToUpper(s string) string {
	var result []rune
	for _, i := range s {
		if i >= 'a' && i <= 'z' {
			result = append(result, i-32)
		} else {
			result = append(result, i)
		}
	}
	return string(result)
}
