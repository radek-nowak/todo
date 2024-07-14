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

type Table interface {
	NumColumns() int
	NumRows() int
	ColumnName(index int) string
	Value(row, column int) string
}

type Column interface {
	Name() string
	MaxLen() int
}

func Display(tl todo.TodoList) {
	todos := tl.GetTodos()
	idLength, taskLength := getMaxLengths(todos)

	// Header
	headerTop, headerMiddle, headerBottom := headers2(idLength, taskLength)

	fmt.Println(headerTop)
	fmt.Println(headerMiddle)
	fmt.Println(headerBottom)

	// Rows
	rowSeparator := rowSeparator(idLength, taskLength)

	// Footer
	// footer := footer(idLength, taskLength)

	for i, todo := range todos {
		fmt.Printf("%s %-*d %s %-*s %s %-5t %s\n", vertical, idLength, todo.ID, vertical, taskLength, todo.Task, vertical, todo.Done, vertical)
		if i < len(todos)-1 { // do not add separator for the las task
			fmt.Println(rowSeparator)
		}
	}

	fmt.Println(headerBottom)
}

func headers(columns []Column) (string, string, string) {

	headerTop:=headerTop(columns)
	headerMiddle:=headerMiddle(columns)
	headerBottom:=headerBottom(columns)

	return headerTop, headerMiddle, headerBottom
}

func headerTop(columns []Column) string {
	var sb strings.Builder

	sb.WriteString(topLeft)

	// move to the caller
	if len(columns) == 0 {
		panic("No columns provided")
	}

	if len(columns) == 1 {
		sb.WriteString(topRight)
		return sb.String()
	}

	for i := 1; i < len(columns); i++ {
		sb.WriteString(strings.Repeat(horizontal, columns[i-1].MaxLen()+2))
		sb.WriteString(topMiddle)
		sb.WriteString(strings.Repeat(horizontal, columns[i].MaxLen()+2))
	}
	sb.WriteString(topRight)

	return sb.String()
}

func headerBottom(columns []Column) string {
	var sb strings.Builder

	sb.WriteString(bottomLeft)

	// move to the caller
	if len(columns) == 0 {
		panic("No columns provided")
	}

	if len(columns) == 1 {
		sb.WriteString(bottomRight)
		return sb.String()
	}

	for i := 1; i < len(columns); i++ {
		sb.WriteString(strings.Repeat(horizontal, columns[i-1].MaxLen()+2))
		sb.WriteString(bottomMiddle)
		sb.WriteString(strings.Repeat(horizontal, columns[i].MaxLen()+2))
	}
	sb.WriteString(bottomRight)

	return sb.String()
}

func footer(columns []Column) string {
	var sb strings.Builder

	sb.WriteString(bottomLeft)

	// move to the caller
	if len(columns) == 0 {
		panic("No columns provided")
	}

	if len(columns) == 1 {
		sb.WriteString(bottomRight)
		return sb.String()
	}

	for i := 1; i < len(columns); i++ {
		sb.WriteString(strings.Repeat(horizontal, columns[i-1].MaxLen()+2))
		sb.WriteString(bottomMiddle)
		sb.WriteString(strings.Repeat(horizontal, columns[i].MaxLen()+2))
	}
	sb.WriteString(bottomRight)

	return sb.String()
}

func headerMiddle(columns []Column) string {
	var sb strings.Builder

	sb.WriteString(vertical)

	for _, col := range columns {
		name := col.Name()
		padding := col.MaxLen() + 2
		sb.WriteString(" ")
		sb.WriteString(centerText(name, padding-2))
		sb.WriteString(" ")
		sb.WriteString(vertical)
	}
	return sb.String()
}

func centerText(text string, width int) string {
	if len(text) >= width {
		return text
	}
	leftPadding := (width - len(text)) / 2
	rightPadding := width - len(text) - leftPadding
	return strings.Repeat(" ", leftPadding) + text + strings.Repeat(" ", rightPadding)
}

func headers2(idLength int, taskLength int) (string, string, string) {
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
	return "", headerMiddle, headerBottom
}

func rowSeparator()  {
	
}

func rowSeparator2(idLength int, taskLength int) string {
	return fmt.Sprintf(
		"%s%s%s%s%s%s%s",
		middleLeft,
		strings.Repeat(horizontal, idLength+2),
		middle,
		strings.Repeat(horizontal, taskLength+2),
		middle,
		strings.Repeat(horizontal, 7),
		middleRight,
	)
}

func footer2(idLength int, taskLength int) string {
	return fmt.Sprintf(
		"%s%s%s%s%s%s%s",
		bottomLeft,
		strings.Repeat(horizontal, idLength+2),
		bottomMiddle,
		strings.Repeat(horizontal, taskLength+2),
		bottomMiddle,
		strings.Repeat(horizontal, 7),
		bottomRight,
	)
}

func getMaxLengths(todos []todo.Todo) (int, int) {
	maxIdHeaderLen := 2
	maxTaskHeaderLen := 10

	for _, todo := range todos {
		idLen := len(fmt.Sprintf("%d", todo.ID))
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
