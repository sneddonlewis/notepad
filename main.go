package main

import (
	"fmt"
	. "notepad/ui"
)

func main() {
	fmt.Println("Enter the maximum number of notes:")
	var size int
	fmt.Scanln(&size)
	RunUi(size)
}
