package table

import "strings"

type lineType int

const (
	LINE_INNER lineType = iota
	LINE_TOP
	LINE_SUBTOP
	LINE_BOTTOM
)

type Separator struct {
	where lineType
}

func (s *Separator) Render(style *renderStyle) string {
	parts := []string{}
	for i := 0; i < style.columns; i++ {
		w := style.PaddingLeft + style.CellWidth(i) + style.PaddingRight
		parts = append(parts, strings.Repeat(style.BorderX, w))
	}

	switch s.where {
	case LINE_TOP:
		return style.BorderTopLeft + strings.Join(parts, style.BorderTop) + style.BorderTopRight
	case LINE_SUBTOP:
		return style.BorderLeft + strings.Join(parts, style.BorderTop) + style.BorderRight
	case LINE_BOTTOM:
		return style.BorderBottomLeft + strings.Join(parts, style.BorderBottom) + style.BorderBottomRight
	case LINE_INNER:
		return style.BorderLeft + strings.Join(parts, style.BorderI) + style.BorderRight
	}
	panic("not reached")
}
