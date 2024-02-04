package piscine

func SplitWhiteSpaces(s string) []string {
	var sent []string
	word := -1
	for i, char := range s {
		if char == ' ' || char == '\t' || char == '\n' {
			if word != -1 {
				sent = append(sent, s[word:i])
				word = -1
			}
		} else if word == -1 {
			word = 1
		}
		if word != -1 {
			sent = append(sent, s[word:])
		}

	}
	return sent
}
