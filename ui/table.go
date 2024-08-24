package ui

import (
	"fmt"
	"strings"
)

type Displayable interface {
	Schema() []string
	Data() [][]any
	ColumnWidths() []int
}

// Table borders
const (
	topLeft      = "╭"
	topRight     = "╮"
	bottomLeft   = "╰"
	bottomRight  = "╯"
	horizontal   = "─"
	vertical     = "│"
	middleLeft   = "├"
	middleRight  = "┤"
	middle       = "┼"
	topMiddle    = "┬"
	bottomMiddle = "┴"
)

func Display(d Displayable) {
	schema := d.Schema()
	data := d.Data()
	colWidths := d.ColumnWidths()

	// Display the header
	headerTop := formatElement(colWidths, topLeft, topMiddle, topRight)
	headrBottom := formatElement(colWidths, middleLeft, middle, middleRight)
	headerMiddle := formatHeaderMiddle(schema, colWidths)
	fmt.Println(headerTop)
	fmt.Println(headerMiddle)
	fmt.Println(headrBottom)

	// Display the rows
	for i, row := range data {
		wrappedRows := wrapRow(row, colWidths)
		for j, wrappedRow := range wrappedRows {
			if j == 0 {
				for k, col := range wrappedRow {
					fmt.Printf("%s %-*s ", vertical, colWidths[k], col)
				}
				fmt.Println(vertical)
			} else {
				for k, col := range wrappedRow {
					fmt.Printf("%s %-*s ", vertical, colWidths[k], col)
				}
				fmt.Println(vertical)
			}
		}
		if i < len(data)-1 {
			fmt.Println(formatRowSeparator(colWidths))
		}
	}

	// Display the footer
	footer := formatElement(colWidths, bottomLeft, bottomMiddle, bottomRight)
	fmt.Println(footer)
}

func formatElement(colWidths []int, left string, middle string, right string) string {
	var uiElement strings.Builder
	uiElement.WriteString(left)
	for i, width := range colWidths {
		if i > 0 {
			uiElement.WriteString(middle)
		}
		uiElement.WriteString(strings.Repeat(horizontal, width+2))
	}
	uiElement.WriteString(right)
	return uiElement.String()
}

func formatHeaderMiddle(schema []string, colWidths []int) string {
	var headerMiddle strings.Builder
	headerMiddle.WriteString(vertical)
	for i, col := range schema {
		headerMiddle.WriteString(fmt.Sprintf(" %-*s %s", colWidths[i], col, vertical))
	}
	return headerMiddle.String()
}

func formatRowSeparator(colWidths []int) string {
	var rowSeparator strings.Builder
	rowSeparator.WriteString(middleLeft)
	for i, width := range colWidths {
		if i > 0 {
			rowSeparator.WriteString(middle)
		}
		rowSeparator.WriteString(strings.Repeat(horizontal, width+2))
	}
	rowSeparator.WriteString(middleRight)
	return rowSeparator.String()
}

func formatFooter(colWidths []int) string {
	var footer strings.Builder
	footer.WriteString(bottomLeft)
	for i, width := range colWidths {
		if i > 0 {
			footer.WriteString(bottomMiddle)
		}
		footer.WriteString(strings.Repeat(horizontal, width+2))
	}
	footer.WriteString(bottomRight)
	return footer.String()
}

func wrapRow(row []interface{}, colWidths []int) [][]string {
	var wrappedRows [][]string
	maxLines := 1

	for i, col := range row {
		colStr := fmt.Sprintf("%v", col)
		wrapped := wrapText(colStr, colWidths[i])
		if len(wrapped) > maxLines {
			maxLines = len(wrapped)
		}
		wrappedRows = append(wrappedRows, wrapped)
	}

	result := make([][]string, maxLines)
	for i := range result {
		result[i] = make([]string, len(row))
		for j := range row {
			if i < len(wrappedRows[j]) {
				result[i][j] = wrappedRows[j][i]
			} else {
				result[i][j] = ""
			}
		}
	}

	return result
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
