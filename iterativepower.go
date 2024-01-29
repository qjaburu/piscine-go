package main

func iterativePower(nb int, power int) int {
	result := 1
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	for power > 1 {
		result *= nb
		power--
	}
	return result
}
