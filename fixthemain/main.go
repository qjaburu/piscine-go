package main

import "fmt"

type Door struct {
	state int
}

const (
	OPEN  = 1
	CLOSE = 0
)

func PrintStr(s string) {
	for _, r := range s {
		fmt.Printf("%c", r)
	}
	fmt.Println()
}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...")
	ptrDoor.state = OPEN
}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...")
	ptrDoor.state = CLOSE
}

func IsDoorOpen(Door Door) bool {
	PrintStr("Is the Door opened?")
	return Door.state == OPEN
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("Is the Door closed?")
	return ptrDoor.state == CLOSE
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(*door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}
