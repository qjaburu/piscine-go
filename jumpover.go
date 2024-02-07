package piscine

func JumpOver(str string) string {
	var value string
	for i, char := range str {
		if (i+1)%3 == 0 {
			value += string(char)
		}
	}
	if value == "" {
		return "\n"
	}
	return value + "\n"
}
