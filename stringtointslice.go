package piscine

func StringToIntSlice(str string) []int {
	var changed []int
	for _, char := range str {
		changed = append(changed, int(char))
	}
	return changed
}
