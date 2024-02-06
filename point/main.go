package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func PrintStr(s string) {
	arrayStr := []rune(s)
	for i := range arrayStr {
		z01.PrintRune(arrayStr[i])
	}
}

func PrintValue(n int) {
	v := '0'
	if n/10 > 0 {
		PrintValue(n / 10)
	}
	for i := 0; i < n%10; i++ {
		v++
	}
	z01.PrintRune(v)
}

func main() {
	points := &point{}
	setPoint(points)
	PrintStr(" x=  ")
	PrintValue(points.x)
	PrintStr(", y = ")
	PrintValue(points.y)
	z01.PrintRune('\n')
}
