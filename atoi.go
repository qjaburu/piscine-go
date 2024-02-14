package main

import "fmt"

func Atoi(s string) int {
	sign := 1
	result := 0
	for i, c := range s {
		if i == 0 && c == '+' {
			sign = 1
		}
		if i == 0 && c == '-' {
			sign = -1
		}
		if i > 0 && c == '-' {
			return 0
		}
		if i > 0 && c == '+' {
			return 0
		}
		if c >= '0' && c <= '9' {
			digit := int(c - '0')
			result = result*10 + digit
		}
		result = result * sign

	}
	return result
}

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}
