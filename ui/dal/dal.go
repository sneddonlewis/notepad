package dal

import (
	"fmt"
	"strconv"
	"strings"
)

type Notepad struct {
	notes    []string
	size     int
	capacity int
}

func NewNotepad(capacity int) *Notepad {
	notes := make([]string, capacity)
	notepad := Notepad{
		notes:    notes,
		size:     0,
		capacity: capacity,
	}
	return &notepad
}

func (notepad *Notepad) Create(entry string) string {
	if entry == "create" {
		return "[Error] Missing note argument"
	}
	if notepad.size == notepad.capacity {
		return "[Error] Notepad is full"
	}
	notepad.notes[notepad.size] = entry
	notepad.size += 1
	return "[Ok] The note was successfully created"
}

func (notepad *Notepad) List() {
	if notepad.size == 0 {
		fmt.Println("[Info] Notepad is empty")
		return
	}
	for index, note := range notepad.notes {
		if notepad.size == index {
			break
		}
		fmt.Printf("[Info] %d: %s\n", index+1, note)
	}
}

func (notepad *Notepad) Clear() string {
	notepad.size = 0
	return "[Ok] All notes were successfully deleted"
}

func (notepad *Notepad) Update(args string) string {
	if args == "update" {
		return "[Error] Missing position argument"
	}
	sArgs := strings.Split(args, " ")
	posAsStr := sArgs[0]
	position, err := strconv.Atoi(posAsStr)
	if err != nil {
		return fmt.Sprintf("[Error] Invalid position: %s", posAsStr)
	}
	if len(sArgs) < 2 {
		return "[Error] Missing note argument"
	}
	if position <= 0 || position > notepad.capacity {
		return fmt.Sprintf(
			"[Error] Position %d is out of the boundary [1, %d]",
			position,
			notepad.capacity)
	}

	var sb strings.Builder
	for i, word := range sArgs {
		if i == 0 {
			continue
		}
		if i != 1 {
			sb.WriteString(" ")
		}
		sb.WriteString(word)
	}

	if position > notepad.size {
		return "[Error] There is nothing to update"
	}

	notepad.notes[position-1] = sb.String()
	return fmt.Sprintf(
		"[Ok] The note at position %d was successfully updated",
		position)
}

func (notepad *Notepad) Delete(args string) string {
	if args == "delete" {
		return "[Error] Missing position argument"
	}
	position, err := strconv.Atoi(args)
	if err != nil {
		return fmt.Sprintf("[Error] Invalid position: %s", args)
	}
	if position <= 0 || position > notepad.capacity {
		return fmt.Sprintf(
			"[Error] Position %d is out of the boundary [1, %d]",
			position,
			notepad.capacity)
	}

	if position > notepad.size {
		return "[Error] There is nothing to delete"
	}

	notes := make([]string, notepad.capacity)
	count := 0
	for index, note := range notepad.notes {
		if index == position-1 {
			continue
		}
		notes[count] = note
		count += 1
	}
	notepad.notes = notes
	notepad.size -= 1

	return fmt.Sprintf(
		"[Ok] The note at position %d was successfully deleted",
		position)
}
