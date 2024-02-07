package piscine

func Abort(a, b, c, d, e int) int {
	store := []int{a, b, c, d, e}
	for i := 0; i < len(store); i++ {
		for j := i + 1; j < len(store); j++ {
			if store[j] < store[i] {
				store[i], store[j] = store[j], store[i]
			}
		}
	}
	return store[2]
}
