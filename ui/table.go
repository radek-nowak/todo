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

const (
	idLength   = 3
	taskLength = 55
	doneLength = 5
)

func Display(tl todo.TodoList) {
	todos := tl.GetTodos()

	// Display the header
	fmt.Println(formatHeaderTop())
	fmt.Println(formatHeaderMiddle())
	fmt.Println(formatHeaderBottom())

	// Display the rows
	for i, todo := range todos {
		taskLines := wrapText(todo.Task, taskLength)
		for j, line := range taskLines {
			if j == 0 {
				fmt.Printf("%s %-*d %s %-*s %s %-*t %s\n", vertical, idLength, todo.ID, vertical, taskLength, line, vertical, doneLength, todo.Done, vertical)
			} else {
				fmt.Printf("%s %-*s %s %-*s %s %-*s %s\n", vertical, idLength, "", vertical, taskLength, line, vertical, doneLength, "", vertical)
			}
		}
		if i < len(todos)-1 {
			fmt.Println(formatRowSeparator())
		}
	}

	// Display the footer
	fmt.Println(formatFooter())
}

func formatHeaderTop() string {
	return fmt.Sprintf(
		"%s%s%s%s%s%s%s",
		topLeft,
		strings.Repeat(horizontal, idLength+2),
		topMiddle,
		strings.Repeat(horizontal, taskLength+2),
		topMiddle,
		strings.Repeat(horizontal, doneLength+2),
		topRight,
	)
}

func formatHeaderMiddle() string {
	return fmt.Sprintf(
		"%s %-*s %s %-*s %s %-*s %s",
		vertical,
		idLength,
		"ID",
		vertical,
		taskLength,
		"Task",
		vertical,
		doneLength,
		"Done",
		vertical,
	)
}

func formatHeaderBottom() string {
	return fmt.Sprintf(
		"%s%s%s%s%s%s%s",
		middleLeft,
		strings.Repeat(horizontal, idLength+2),
		middle,
		strings.Repeat(horizontal, taskLength+2),
		middle,
		strings.Repeat(horizontal, doneLength+2),
		middleRight,
	)
}

func formatRowSeparator() string {
	return fmt.Sprintf(
		"%s%s%s%s%s%s%s",
		middleLeft,
		strings.Repeat(horizontal, idLength+2),
		middle,
		strings.Repeat(horizontal, taskLength+2),
		middle,
		strings.Repeat(horizontal, doneLength+2),
		middleRight,
	)
}

func formatFooter() string {
	return fmt.Sprintf(
		"%s%s%s%s%s%s%s",
		bottomLeft,
		strings.Repeat(horizontal, idLength+2),
		bottomMiddle,
		strings.Repeat(horizontal, taskLength+2),
		bottomMiddle,
		strings.Repeat(horizontal, doneLength+2),
		bottomRight,
	)
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
