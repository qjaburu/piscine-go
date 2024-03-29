package piscine

func Atoi(s string) int {
	sign := 1
	result := 0
	for i, c := range s {
		if i == 0 && c == '+' {
			sign = 1
			continue
		}
		if i == 0 && c == '-' {
			sign = -1
			continue
		}
		if i > 0 && c == '-' {
			return 0
		}
		if i > 0 && c == '+' {
			return 0
		}
		if i > 0 && c == ' ' {
			return 0
		}
		if c >= '0' && c <= '9' {
			digit := int(c - '0')
			result = result*10 + digit
		} else {
			return 0
		}

	}
	return result * sign
}
