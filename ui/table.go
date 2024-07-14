package ui

import (
	"fmt"
	todo "go_todo/todo/model"
	"strings"
)

const (
	topLeft      = "╔"
	topRight     = "╗"
	bottomLeft   = "╚"
	bottomRight  = "╝"
	horizontal   = "═"
	vertical     = "║"
	middleLeft   = "╠"
	middleRight  = "╣"
	middle       = "╬"
	topMiddle    = "╦"
	bottomMiddle = "╩"
)

func Display(tl todo.TodoList) {
	todos := tl.GetTodos()
	idLength, taskLength := 3, 55

	// Header
	headerTop := fmt.Sprintf(
		"%s%s%s%s%s%s%s",
		topLeft,
		strings.Repeat(horizontal, idLength+2),
		topMiddle,
		strings.Repeat(horizontal, taskLength+2),
		topMiddle,
		strings.Repeat(horizontal, 7),
		topRight,
	)

	headerMiddle := fmt.Sprintf(
		"%s %-*s %s %-*s %s %-5s %s",
		vertical,
		idLength,
		"ID",
		vertical,
		taskLength,
		"Task",
		vertical,
		"Done",
		vertical,
	)

	headerBottom := fmt.Sprintf(
		"%s%s%s%s%s%s%s",
		middleLeft,
		strings.Repeat(horizontal, idLength+2),
		middle,
		strings.Repeat(horizontal, taskLength+2),
		middle,
		strings.Repeat(horizontal, 7),
		middleRight,
	)

	fmt.Println(headerTop)
	fmt.Println(headerMiddle)
	fmt.Println(headerBottom)

	// Rows
	rowSeparator := fmt.Sprintf(
		"%s%s%s%s%s%s%s",
		middleLeft,
		strings.Repeat(horizontal, idLength+2),
		middle,
		strings.Repeat(horizontal, taskLength+2),
		middle,
		strings.Repeat(horizontal, 7),
		middleRight,
	)

	// Footer
	footer := fmt.Sprintf(
		"%s%s%s%s%s%s%s",
		bottomLeft,
		strings.Repeat(horizontal, idLength+2),
		bottomMiddle,
		strings.Repeat(horizontal, taskLength+2),
		bottomMiddle,
		strings.Repeat(horizontal, 7),
		bottomRight,
	)

	for i, todo := range todos {
		taskLines := wrapText(todo.Task, taskLength)
		for j, line := range taskLines {
			if j == 0 {
				fmt.Printf("%s %-*d %s %-*s %s %-5t %s\n", vertical, idLength, todo.ID, vertical, taskLength, line, vertical, todo.Done, vertical)
			} else {
				fmt.Printf("%s %-*s %s %-*s %s %-5s %s\n", vertical, idLength, "", vertical, taskLength, line, vertical, "", vertical)
			}
		}
		if i < len(todos)-1 {
			fmt.Println(rowSeparator)
		}
	}

	fmt.Println(footer)
}

func wrapText(text string, length int) []string {
	if len(text) <= length {
		return []string{text}
	}

	var wrapped []string
	for len(text) > length {
		wrapped = append(wrapped, text[:length])
		text = text[length:]
	}
	wrapped = append(wrapped, text)

	return wrapped
}

