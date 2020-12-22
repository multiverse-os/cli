package graphics

import (
	"strconv"
)

type GraphicsMode interface {
	getGraphicsModeString() string
	getGraphicsModeInt() int
	addColorOffset(Color) string
}

type graphicsMode int

func (g graphicsMode) getGraphicsModeString() string {
	return strconv.Itoa(int(g))
}

func (g graphicsMode) getGraphicsModeInt() int {
	return int(g)
}

func (g graphicsMode) addColorOffset(c Color) string {
	return strconv.Itoa(int(g) + c.getColorInt())
}

const (
	setGraphics string = "\x1b[%sm"

	ResetGraphics          graphicsMode = 0
	Bold                   graphicsMode = 1 // increase brightness
	Faded                  graphicsMode = 2 // lower brightness
	Italics                graphicsMode = 3
	Underlined             graphicsMode = 4
	SlowFlashing           graphicsMode = 5
	FastFlashing           graphicsMode = 6
	Negative               graphicsMode = 7
	Hidden                 graphicsMode = 8
	CrossedOut             graphicsMode = 9
	DefaultFont            graphicsMode = 10
	AlternativeFont        graphicsMode = 11 // plus 1 - 2 alt font - not checked
	BoldAndUnderlined      graphicsMode = 21
	ResetBoldAndBrightness graphicsMode = 22
	ResetItalics           graphicsMode = 23
	ResetUnderlined        graphicsMode = 24
	ResetFlashing          graphicsMode = 25
	ResetNegative          graphicsMode = 27
	ResetHidden            graphicsMode = 28
	ResetCrossedOut        graphicsMode = 29
	setForegroundColor     graphicsMode = 30
	ResetForegroundColor   graphicsMode = 39
	setBackgroundColor     graphicsMode = 40
	ResetBackgroundColor   graphicsMode = 49
)

type Color interface {
	getColorInt() int
	getColorString() string
}

type color int

func (c color) getColorInt() int {
	return int(c)
}

func (c color) getColorString() string {
	return strconv.Itoa(c.getColorInt())
}

const (
	Black   color = 0
	Red     color = 1
	Green   color = 2
	Yellow  color = 3
	Blue    color = 4
	Magenta color = 5
	Cyan    color = 6
	White   color = 7
	NoColor color = 8
)

type ClearMode interface {
	getModeInt() int
}

type clearMode int

func (c clearMode) getModeInt() int {
	return int(c)
}

const (
	clearScreen string = "\x1b[%dJ"

	ClearAfterCursor  clearMode = 0 // 0 - clears all from cursors position to the end of terminal
	ClearBeforeCursor clearMode = 1 // 1 - clears all from cursors position to the start of terminal
	ClearAll          clearMode = 2 // 2 - clears all
)

type ClearLineMode interface {
	getLineModeInt() int
}

type clearLineMode int

func (c clearLineMode) getLineModeInt() int {
	return int(c)
}

const (
	clearLine string = "\x1b[%dK"

	ClearLineAfterCursor  clearLineMode = 0 // 0 - clears all from cursors position to the end of Line
	ClearLineBeforeCursor clearLineMode = 1 // 1 - clears all from cursors position to the start of line
	ClearAllLine          clearLineMode = 2 // 2 - clears all line
)
