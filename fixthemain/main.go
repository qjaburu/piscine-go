package piscine

import "fmt"

type Door struct {
	state int
}

func OpenDoor(ptrDoor *Door) {
	fmt.Println("Door Opening...")
	ptrDoor.state = 1
}

func CloseDoor(ptrDoor *Door) {
	fmt.Println("Door Closing...")
	ptrDoor.state = 0
}

func IsDoorOpen(ptrDoor *Door) bool {
	fmt.Println("Is the Door opened ?")
	return ptrDoor.state == 1
}

func IsDoorClose(ptrDoor *Door) bool {
	fmt.Println("Is the Door closed ?")
	return ptrDoor.state == 0
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == 1 {
		CloseDoor(door)
	}
}
