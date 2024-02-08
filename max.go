package piscine

func Max(s []int) int {
	if len(s) == 0 {
		return 0
	}
	max := s[0]
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	return max
}
