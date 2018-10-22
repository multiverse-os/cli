package table

import (
	"strings"
)

type Row struct {
	cells []*Cell
}

func CreateRow(items []interface{}) *Row {
	row := &Row{cells: []*Cell{}}
	for _, item := range items {
		row.AddCell(item)
	}
	return row
}

func (r *Row) AddCell(item interface{}) {
	if c, ok := item.(*Cell); ok {
		c.column = len(r.cells)
		r.cells = append(r.cells, c)
	} else {
		r.cells = append(r.cells, createCell(len(r.cells), item, nil))
	}
}

func (r *Row) Render(style *renderStyle) string {
	renderedCells := []string{}
	for _, c := range r.cells {
		renderedCells = append(renderedCells, c.Render(style))
	}
	return style.BorderY + strings.Join(renderedCells, style.BorderY) + style.BorderY
}
