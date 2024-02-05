package main

import "fmt"

func Capitalize(s string) string {
	var result string
	var next bool = true

	for _, char := range s {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char >= '0' && char <= '9' {
			if next {
				if char >= 'a' && char <= 'z' {
					result += string(char - 32)
				} else {
					result += string(char)
				}
				next = false
			} else {
				if char >= 'A' && char <= 'Z' {
					result += string(char + 32)
				} else {
					result += string(char)
				}
			}
		} else {
			result += string(char)
			next = true
		}
	}
	return result
}

func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}
