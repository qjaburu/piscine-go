package piscine

func Capitalize(s string) string {
	var result string
	var next bool
	for _, n := range s {
		if (n >= 'a' && n <= 'z') || (n >= 'A' && n <= 'z') || (n >= '0' && n <= '9') {
			if next {
				if n >= 'a' && n <= 'z' {
					result += string(n - ('a' - 'A'))
				} else {
					result += string(n)
				}
				next = false

			} else {
				if n >= 'A' && n <= 'Z' {
					result += string(n + ('a' - 'A'))
				} else {
					result += string(n)
				}
			}
		} else {
			result += string(n)
			next = true
		}
	}
	return result
}
