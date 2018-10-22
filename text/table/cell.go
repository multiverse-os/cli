package table

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (
	colorFilter = regexp.MustCompile(`\033\[(?:\d+(?:;\d+)*)?m`)
)

type Cell struct {
	column         int
	formattedValue string
	alignment      *tableAlignment
	colSpan        int
}

func CreateCell(v interface{}, style *CellStyle) *Cell {
	return createCell(0, v, style)
}

func createCell(column int, v interface{}, style *CellStyle) *Cell {
	cell := &Cell{column: column, formattedValue: renderValue(v), colSpan: 1}
	if style != nil {
		cell.alignment = &style.Alignment
		if style.ColSpan != 0 {
			cell.colSpan = style.ColSpan
		}
	}
	return cell
}

func (c *Cell) Width() int {
	return runewidth.StringWidth(filterColorCodes(c.formattedValue))
}

func filterColorCodes(s string) string {
	return colorFilter.ReplaceAllString(s, "")
}

func (c *Cell) Render(style *renderStyle) (buffer string) {
	if c.alignment == nil {
		c.alignment = &style.Alignment
	}

	// left padding
	buffer += strings.Repeat(" ", style.PaddingLeft)

	// append the main value and handle alignment
	buffer += c.alignCell(style)

	// right padding
	buffer += strings.Repeat(" ", style.PaddingRight)

	// this handles escaping for, eg, Markdown, where we don't care about the
	// alignment quite as much
	if style.replaceContent != nil {
		buffer = style.replaceContent(buffer)
	}

	return buffer
}

func (c *Cell) alignCell(style *renderStyle) string {
	buffer := ""
	width := style.CellWidth(c.column)

	if c.colSpan > 1 {
		for i := 1; i < c.colSpan; i++ {
			w := style.CellWidth(c.column + i)
			if w == 0 {
				break
			}
			width += style.PaddingLeft + w + style.PaddingRight + utf8.RuneCountInString(style.BorderY)
		}
	}

	switch *c.alignment {

	default:
		buffer += c.formattedValue
		if l := width - c.Width(); l > 0 {
			buffer += strings.Repeat(" ", l)
		}

	case AlignLeft:
		buffer += c.formattedValue
		if l := width - c.Width(); l > 0 {
			buffer += strings.Repeat(" ", l)
		}

	case AlignRight:
		if l := width - c.Width(); l > 0 {
			buffer += strings.Repeat(" ", l)
		}
		buffer += c.formattedValue

	case AlignCenter:
		left, right := 0, 0
		if l := width - c.Width(); l > 0 {
			lf := float64(l)
			left = int(math.Floor(lf / 2))
			right = int(math.Ceil(lf / 2))
		}
		buffer += strings.Repeat(" ", left)
		buffer += c.formattedValue
		buffer += strings.Repeat(" ", right)
	}

	return buffer
}

// Format the raw value as a string depending on the type
func renderValue(v interface{}) string {
	switch vv := v.(type) {
	case string:
		return vv
	case bool:
		return strconv.FormatBool(vv)
	case int:
		return strconv.Itoa(vv)
	case int64:
		return strconv.FormatInt(vv, 10)
	case uint64:
		return strconv.FormatUint(vv, 10)
	case float64:
		return strconv.FormatFloat(vv, 'f', 2, 64)
	case fmt.Stringer:
		return vv.String()
	}
	return fmt.Sprintf("%v", v)
}
