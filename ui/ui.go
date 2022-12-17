package ui

import (
	"bufio"
	"fmt"
	. "notepad/ui/dal"
	"os"
	"strings"
)

const prompt = "Enter a command and data:"

func RunUi(size int) {
	notes := NewNotepad(size)
	for {
		fmt.Println(prompt)
		cmd := parseInput()
		cmd.Execute(notes)
	}
}

type Cmd struct {
	command, data string
}

func parseInput() Cmd {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	line := scanner.Text()
	words := strings.Split(line, " ")
	cmd := words[0]
	data := strings.TrimPrefix(strings.Join(words, " "), cmd+" ")
	return Cmd{cmd, data}
}

func (cmd *Cmd) Execute(notes *Notepad) {
	switch cmd.command {
	case "exit":
		fmt.Println("[Info] Bye!")
		os.Exit(1)
	case "create":
		handleCreate(notes, cmd.data)
	case "list":
		handleList(notes)
	case "clear":
		handleClear(notes)
	case "update":
		handleUpdate(notes, cmd.data)
	case "delete":
		handleDelete(notes, cmd.data)
	default:
		fmt.Println("[Error] Unknown command")
	}
}

func handleCreate(notes *Notepad, input string) {
	fmt.Println(notes.Create(input))
}

func handleList(notes *Notepad) {
	notes.List()
}

func handleClear(notes *Notepad) {
	fmt.Println(notes.Clear())
}

func handleUpdate(notes *Notepad, input string) {
	fmt.Println(notes.Update(input))
}

func handleDelete(notes *Notepad, input string) {
	fmt.Println(notes.Delete(input))
}
