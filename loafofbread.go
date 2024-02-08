package piscine

func LoafOfBread(str string) string {
	if str == "" {
		return "\n"
	}
	if len(str) < 5 {
		return "invalid output"
	}
	result := ""
	counter := 0
	for i, char := range str {
		if char != ' ' && counter != 5 {
			result += string(char)
			counter++
		} else if counter == 5 {
			result += " "
			counter = 0
		}
		if i == len(str)-1 && len(result) > 0 && result[len(result)-1] == ' ' {
			result = result[:len(result)-1]
		}
	}
	result += "\n"
	return result
}
