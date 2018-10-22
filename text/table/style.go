package table

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type tableAlignment int

const (
	AlignLeft   = tableAlignment(1)
	AlignCenter = tableAlignment(2)
	AlignRight  = tableAlignment(3)
)

type TableStyle struct {
	SkipBorder        bool
	BorderX           string
	BorderY           string
	BorderI           string
	BorderTop         string
	BorderBottom      string
	BorderRight       string
	BorderLeft        string
	BorderTopLeft     string
	BorderTopRight    string
	BorderBottomLeft  string
	BorderBottomRight string
	PaddingLeft       int
	PaddingRight      int
	Width             int
	Alignment         tableAlignment
	htmlRules         htmlStyleRules
}

type CellStyle struct {
	Alignment tableAlignment
	ColSpan   int
}

var DefaultStyle = &TableStyle{
	SkipBorder: false,
	BorderX:    "-", BorderY: "|", BorderI: "+",
	PaddingLeft: 1, PaddingRight: 1,
	Width:     80,
	Alignment: AlignLeft,

	// FIXME: the use of a Width here may interact poorly with a changing
	// MaxColumns value; we don't set MaxColumns here because the evaluation
	// order of a var and an init value adds undesired subtlety.
}

type renderStyle struct {
	cellWidths     map[int]int
	columns        int
	replaceContent func(string) string
	TableStyle
}

func (s *TableStyle) setUtfBoxStyle() {
	s.BorderX = "─"
	s.BorderY = "│"
	s.BorderI = "┼"
	s.BorderTop = "┬"
	s.BorderBottom = "┴"
	s.BorderLeft = "├"
	s.BorderRight = "┤"
	s.BorderTopLeft = "╭"
	s.BorderTopRight = "╮"
	s.BorderBottomLeft = "╰"
	s.BorderBottomRight = "╯"
}

// setAsciiBoxStyle changes the border characters back to their defaults
func (s *TableStyle) setAsciiBoxStyle() {
	s.BorderX = "-"
	s.BorderY = "|"
	s.BorderI = "+"
	s.BorderTop, s.BorderBottom, s.BorderLeft, s.BorderRight = "", "", "", ""
	s.BorderTopLeft, s.BorderTopRight, s.BorderBottomLeft, s.BorderBottomRight = "", "", "", ""
	s.fillStyleRules()
}

// fillStyleRules populates members of the TableStyle box-drawing specification
// with BorderI as the default.
func (s *TableStyle) fillStyleRules() {
	if s.BorderTop == "" {
		s.BorderTop = s.BorderI
	}
	if s.BorderBottom == "" {
		s.BorderBottom = s.BorderI
	}
	if s.BorderLeft == "" {
		s.BorderLeft = s.BorderI
	}
	if s.BorderRight == "" {
		s.BorderRight = s.BorderI
	}
	if s.BorderTopLeft == "" {
		s.BorderTopLeft = s.BorderI
	}
	if s.BorderTopRight == "" {
		s.BorderTopRight = s.BorderI
	}
	if s.BorderBottomLeft == "" {
		s.BorderBottomLeft = s.BorderI
	}
	if s.BorderBottomRight == "" {
		s.BorderBottomRight = s.BorderI
	}
}

func createRenderStyle(table *Table) *renderStyle {
	style := &renderStyle{TableStyle: *table.Style, cellWidths: map[int]int{}}
	style.TableStyle.fillStyleRules()

	if table.outputMode == outputMarkdown {
		style.buildReplaceContent(table.Style.BorderY)
	}

	// FIXME: handle actually defined width condition

	// loop over the rows and cells to calculate widths
	for _, element := range table.elements {
		// skip separators
		if _, ok := element.(*Separator); ok {
			continue
		}

		// iterate over cells
		if row, ok := element.(*Row); ok {
			for i, cell := range row.cells {
				// FIXME: need to support sizing with colspan handling
				if cell.colSpan > 1 {
					continue
				}
				if style.cellWidths[i] < cell.Width() {
					style.cellWidths[i] = cell.Width()
				}
			}
		}
	}
	style.columns = len(style.cellWidths)

	// calculate actual width
	width := utf8.RuneCountInString(style.BorderLeft) // start at '1' for left border
	internalBorderWidth := utf8.RuneCountInString(style.BorderI)

	lastIndex := 0
	for i, v := range style.cellWidths {
		width += v + style.PaddingLeft + style.PaddingRight + internalBorderWidth
		if i > lastIndex {
			lastIndex = i
		}
	}
	if internalBorderWidth != utf8.RuneCountInString(style.BorderRight) {
		width += utf8.RuneCountInString(style.BorderRight) - internalBorderWidth
	}

	if table.titleCell != nil {
		titleMinWidth := 0 +
			table.titleCell.Width() +
			utf8.RuneCountInString(style.BorderLeft) +
			utf8.RuneCountInString(style.BorderRight) +
			style.PaddingLeft +
			style.PaddingRight

		if width < titleMinWidth {
			// minWidth must be set to include padding of the title, as required
			style.cellWidths[lastIndex] += (titleMinWidth - width)
			width = titleMinWidth
		}
	}
	style.Width = width
	return style
}

func (s *renderStyle) CellWidth(i int) int {
	return s.cellWidths[i]
}

func (s *renderStyle) buildReplaceContent(bad string) {
	replacement := fmt.Sprintf("&#x%02x;", bad)
	s.replaceContent = func(old string) string {
		return strings.Replace(old, bad, replacement, -1)
	}
}
