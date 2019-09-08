package matrix

import (
	"strings"
)

type row map[int]string
type column map[int]row

// A Matrix contains strings addressed by x and y coordinates. They are mutated using matrix.Put(x, y, text).
type Matrix column

// Put sets the given string (text) at the desired x and y coordinate in the matrix. Thereby it can be used to add and
// overwrite data. To remove data, overwrite it with whitespace.
func (m Matrix) Put(x, y int, text string) {
	for i, s := range strings.Split(text, "") {
		if m[y] == nil {
			m[y] = row{}
		}
		m[y][x+i] = s
	}
}

// String renders the the data as a matrix. Empty addresses are filled with whitespace.
func (m Matrix) String() string {
	table := ""

	var totalColumns int
	for i := range m {
		if i > totalColumns {
			totalColumns = i
		}
	}

	for i := 0; i <= totalColumns; i++ {
		row := m[i]
		var buf string

		var totalRows int
		for i := range row {
			if i > totalRows {
				totalRows = i
			}
		}
		for i := 0; i <= totalRows; i++ {
			val := row[i]
			if val == "" {
				val = " "
			}
			buf = buf + val
		}
		table = table + buf + "\n"

	}
	return strings.TrimRight(table, "\n")
}
