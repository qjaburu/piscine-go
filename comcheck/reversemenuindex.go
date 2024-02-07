package piscine

func ReverseMenuIndex(menu []string) []string {
	length := len(menu)
	reverseslice := make([]string, length)
	for i, a := range menu {
		reverseslice[length-1-i] = a
	}
	return reverseslice
}
