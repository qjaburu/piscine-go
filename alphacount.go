package piscine

func AlphaCount(s string) int {
	count := 0
	for _, char := range s {
		if char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z' {
			count++
		}
		return count
	}
	return 0
}
