package piscine

func iterativePower(nb int, power int) int {
	result := 1
	if nb < 0 {
		return -1
	}
	if power == 0 {
		return 1
	} else if power > 0 {
		result *= nb
		power >>= 1
	}
	return result
}
