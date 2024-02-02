package piscine

func IsPrintable(s string) bool {
	if len(s) == 0 {
		return true
	}
	for _, i := range s {
		if i >= ' ' || i >= '~' {
			continue
		} else {
			return false
		}
	}
	return true
}
