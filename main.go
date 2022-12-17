package main

import (
	"fmt"
	. "notepad/ui"
)

const prompt = "Enter a command and data:"

func main() {
	fmt.Println("Enter the maximum number of notes:")
	var size int
	fmt.Scanln(&size)
	notes := NewNotepad(size)
	for {
		fmt.Println(prompt)
		cmd := ParseInput()
		cmd.Execute(notes)
	}
}
