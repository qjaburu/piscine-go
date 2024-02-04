package piscine

func ConcatParams(args []string) string {
	var value string
	for i, v := range args {
		if i == len(args)-1 {
			value += v
			break
		}
		value += v + "\n"
	}
	return value
}
