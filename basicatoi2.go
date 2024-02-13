package piscine

func BasicAtoi2(s string) int {
	var result int
	for _, i := range s {
		if i >= '0' && i <= '9' {
			result = result*10 + int(i-'0')
		} else {
			return 0
		}
	}
	return result
}
