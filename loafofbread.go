package piscine

func LoafOfBread(s string) string {
	if len(s) < 5 {
		return "Invalid Output\n"
	}
	var res string
	for i := 0; i < len(s); i += 6 {
		if i+5 >= len(s) {
			res += s[i:]
			break
		}
		if s[i+5] == ' ' {
			res += s[i:i+5] + " "
			i--
		} else {
			res += s[i : i+5]
		}
	}
	return res
}
