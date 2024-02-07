package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	nsteps := 0
	if start != 1 {
		if start%2 == 0 {
			start = start / 2
		} else {
			start = 3*start + 1
		}
		nsteps++

	}
	return nsteps
}
