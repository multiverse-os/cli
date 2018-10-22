package text

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

// Table containing column titles and rows.
type Table struct {
	titles Row
	rows   []Row
}

// Table row containing cells.
type Row []Cell

// Table cell containing string content, internal padding and alignment.
type Cell struct {
	Content  string
	PadLeft  uint
	PadRight uint
	Align    Alignment
}

// Cell alignment type
type Alignment int

// Cell alignment variants.
const (
	AlignLeft Alignment = iota
	AlignRight
	AlignCenter
)

// Set row that is to be rendered as column titles.
func (table *Table) SetTitles(titles Row) {
	table.titles = titles
}

// Add row to the table.
func (table *Table) AddRow(row Row) {
	table.rows = append(table.rows, row)
}

// Update column widths to that of the largest encountered cell.
func updateColumnWidths(widths *[]int, row Row) {
	n := len(row) - len(*widths)
	for i := 0; i < n; i++ {
		*widths = append(*widths, 0)
	}
	for i, cell := range row {
		l := utf8.RuneCountInString(cell.Content) + int(cell.PadLeft+cell.PadRight)
		if l > (*widths)[i] {
			(*widths)[i] = l
		}
	}
}

// Render table as text.
func (table *Table) Render() []byte {
	var buf bytes.Buffer
	var columnWidths []int

	updateColumnWidths(&columnWidths, table.titles)
	for _, row := range table.rows {
		updateColumnWidths(&columnWidths, row)
	}

	if table.titles != nil {
		fmt.Fprintf(&buf, "%s", table.titles.render(columnWidths))
		// Render the titles underlining.
		for i := 0; i < len(columnWidths); i++ {
			begPad := "-+-"
			if i == 0 {
				begPad = ""
			}
			fmt.Fprintf(&buf, "%s%s", begPad, fill('-', columnWidths[i]))
		}
		fmt.Fprintf(&buf, "\n")
	}
	for _, row := range table.rows {
		fmt.Fprintf(&buf, "%s", row.render(columnWidths))
	}
	return buf.Bytes()
}

// Implements fmt.Stringer
func (table *Table) String() string {
	return string(table.Render())
}

// Render row as text.
func (row *Row) render(columnWidths []int) []byte {
	var buf bytes.Buffer

	for i := 0; i < len(columnWidths); i++ {
		begPad := " | "
		if i == 0 {
			begPad = ""
		}
		fmt.Fprintf(&buf, "%s%s", begPad, row.cellAt(i).render(columnWidths[i]))
	}
	return append(bytes.TrimRight(buf.Bytes(), " "), '\n')
}

// Return cell at the specified column or nil if there is no such column.
func (row *Row) cellAt(column int) *Cell {
	if column >= len(*row) {
		return nil
	}
	return &(*row)[column]
}

// Render cell as text.
// As a special case a cell that is nil is treated as an empty cell.
func (cell *Cell) render(cellWidth int) []byte {
	if cell == nil {
		return fill(' ', cellWidth)
	}
	buf := make([]byte, 0)
	if cell.PadLeft > 0 {
		buf = append(buf, fill(' ', int(cell.PadLeft))...)
	}
	buf = append(buf, cell.Content...)
	if cell.PadRight > 0 {
		buf = append(buf, fill(' ', int(cell.PadRight))...)
	}

	pad := cellWidth - utf8.RuneCount(buf)
	var padLeft, padRight int
	switch cell.Align {
	case AlignRight:
		padLeft = pad
	case AlignCenter:
		// Pad with a bias of more padding to the right.
		padLeft = pad / 2
		padRight = pad - padLeft
	default:
		// Since it's possible to pass values other than the specified
		// constants use AlignLeft as the default.
		padRight = pad
	}
	buf = append(buf, fill(' ', padRight)...)
	return append(fill(' ', padLeft), buf...)
}

// Return a byte slice containing count copies of char.
func fill(char byte, count int) []byte {
	return bytes.Repeat([]byte{char}, count)
}
