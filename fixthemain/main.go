package main

import "github.com/01-edu/z01"

type Door struct {
	state int
}

const (
	OPEN  = 0
	CLOSE = 1
)

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...")
	z01.PrintRune('\n')

	ptrDoor.state = OPEN
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...")
	z01.PrintRune('\n')
	ptrDoor.state = CLOSE
	return true
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door opened ? ")
	z01.PrintRune('\n')

	return ptrDoor.state == OPEN
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ? ")
	z01.PrintRune('\n')

	return ptrDoor.state == CLOSE
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

	if door.state == OPEN {
		CloseDoor(door)
	}
}
