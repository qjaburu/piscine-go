package piscine

func Compact(ptr *[]string) int {
	slice := *ptr
	count := 0

	for i := 0; i < len(slice); {
		if slice[i] != "" {
			count++
			i++
		} else {
			slice = append(slice[:i], slice[i+1:]...)
		}
	}

	*ptr = slice

	return count
}
