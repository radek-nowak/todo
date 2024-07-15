package ui

import (
	"fmt"
	"strings"
)

// Define Displayable interface
type Displayable interface {
	Schema() []string
	Data() [][]interface{}
	ColumnWidths() []int
}

// Define table borders
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

// Display function for Displayable interface
func Display(d Displayable) {
	schema := d.Schema()
	data := d.Data()
	colWidths := d.ColumnWidths()

	// Display the header
	fmt.Println(formatHeaderTop(colWidths))
	fmt.Println(formatHeaderMiddle(schema, colWidths))
	fmt.Println(formatHeaderBottom(colWidths))

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
	fmt.Println(formatFooter(colWidths))
}

func formatHeaderTop(colWidths []int) string {
	var headerTop strings.Builder
	headerTop.WriteString(topLeft)
	for i, width := range colWidths {
		if i > 0 {
			headerTop.WriteString(topMiddle)
		}
		headerTop.WriteString(strings.Repeat(horizontal, width+2))
	}
	headerTop.WriteString(topRight)
	return headerTop.String()
}

func formatHeaderMiddle(schema []string, colWidths []int) string {
	var headerMiddle strings.Builder
	headerMiddle.WriteString(vertical)
	for i, col := range schema {
		headerMiddle.WriteString(fmt.Sprintf(" %-*s %s", colWidths[i], col, vertical))
	}
	return headerMiddle.String()
}

func formatHeaderBottom(colWidths []int) string {
	var headerBottom strings.Builder
	headerBottom.WriteString(middleLeft)
	for i, width := range colWidths {
		if i > 0 {
			headerBottom.WriteString(middle)
		}
		headerBottom.WriteString(strings.Repeat(horizontal, width+2))
	}
	headerBottom.WriteString(middleRight)
	return headerBottom.String()
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


