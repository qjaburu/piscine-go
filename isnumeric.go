package piscine

func IsNumeric(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, q := range s {
		if q >= '0' && q <= '9' { // if you use if condition you have to use "continue"
			continue
		} else {
			return false
		}
	}
	return true
}
