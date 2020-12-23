package termloop

type Text struct {
	x      int
	y      int
	fg     Attr
	bg     Attr
	text   []rune
	canvas []Cell
}

func NewText(x, y int, text string, fg, bg Attr) *Text {
	str := []rune(text)
	c := make([]Cell, len(str))
	for i := range c {
		c[i] = Cell{Ch: str[i], Fg: fg, Bg: bg}
	}
	return &Text{
		x:      x,
		y:      y,
		fg:     fg,
		bg:     bg,
		text:   str,
		canvas: c,
	}
}

func (t *Text) Tick(ev Event)        {}
func (t *Text) Position() (int, int) { return t.x, t.y }
func (t *Text) Size() (int, int)     { return len(t.text), 1 }
func (t *Text) Color() (Attr, Attr)  { return t.fg, t.bg }
func (t *Text) Text() string         { return string(t.text) }

func (t *Text) Draw(s *Screen) {
	w, _ := t.Size()
	for i := 0; i < w; i++ {
		s.RenderCell(t.x+i, t.y, &t.canvas[i])
	}
}

func (t *Text) SetPosition(x, y int) {
	t.x = x
	t.y = y
}

func (t *Text) SetText(text string) {
	t.text = []rune(text)
	c := make([]Cell, len(t.text))
	for i := range c {
		c[i] = Cell{Ch: t.text[i], Fg: t.fg, Bg: t.bg}
	}
	t.canvas = c
}

func (t *Text) SetColor(fg, bg Attr) {
	t.fg = fg
	t.bg = bg
	for i := range t.canvas {
		t.canvas[i].Fg = fg
		t.canvas[i].Bg = bg
	}
}
