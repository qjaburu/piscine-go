package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	list := make(map[string]int)

	var items string

	for _, el := range str { // el is shorr form of element
		if el == 32 {
			list[items] += 1
			items = ""

		} else if el != 32 {
			items += string(byte(el))
		}
	}
	list[items] += 1
	return list
}
