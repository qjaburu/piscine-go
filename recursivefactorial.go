package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 { // this test if this is negative number
		return 0
	} else if nb == 0 || nb == 1 {
		return 1
	} else {
		return nb * RecursiveFactorial(nb-1)
	}
}
