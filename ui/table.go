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
	idLength, taskLength := getMaxLengths(todos)

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
		fmt.Printf("%s %-*d %s %-*s %s %-5t %s\n", vertical, idLength, todo.ID, vertical, taskLength, todo.Task, vertical, todo.Done, vertical)
		if i < len(todos)-1 { // do not add separator for the las task
			fmt.Println(rowSeparator)
		}
	}

	fmt.Println(footer)
}

func getMaxLengths(todos []todo.Todo) (int, int) {
	maxIdHeaderLen := 2
	maxTaskHeaderLen := 10

	for _, todo := range todos {
		idLen := len(string(todo.ID))
		taskLen := len(todo.Task)

		if idLen > maxIdHeaderLen {
			maxIdHeaderLen = idLen
		}

		if taskLen > maxTaskHeaderLen {
			maxTaskHeaderLen = taskLen
		}

	}

	return maxIdHeaderLen, maxTaskHeaderLen
}
