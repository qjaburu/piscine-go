package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || power == 0 {
		return 0
	} else if nb == 0 {
		return 1
	}
	return nb * iterativePower (nb-1 ) )
}
