package piscine

import "fmt"

func IsNumeric(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, q := range s {
		if q >= '0' && q <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}

