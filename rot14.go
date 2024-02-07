package piscine

func Rot14(s string) string {
	result := []rune(s)
	for i, b := range result {
		if b >= 'a' && b <= 'z' {
			result[i] = 'a' + (b-'a'+14)%26
		} else {
			if b >= 'A' && b <= 'Z' {
				result[i] = 'A' + (b-'A'+14)%26
			}
		}
	}
	return string(result)
}
