package graphics

import (
	"fmt"
	"strings"

	"github.com/KlyuchnikovV/cui/cursor"
	"github.com/KlyuchnikovV/cui/low_level/terminal"
	"github.com/KlyuchnikovV/cui/types"
)

type Graphics struct {
	types.ConsoleStream
	cursor.Cursor
	width  int
	height int
}

func New() *Graphics {
	w, h := terminal.GetTerminalSize()
	return &Graphics{
		ConsoleStream: types.NewConsoleStream(),
		Cursor:        *cursor.New(),
		width:         w,
		height:        h,
	}
}

func (g *Graphics) ClearScreen(mode ClearMode) {
	g.Print(fmt.Sprintf(clearScreen, mode.getModeInt()))
}

func (g *Graphics) ClearLine(mode clearLineMode) {
	g.Print(fmt.Sprintf(clearLine, mode.getLineModeInt()))
}

func (g *Graphics) SetGraphics(modes ...GraphicsMode) {
	var result = make([]string, len(modes))
	for i, mode := range modes {
		result[i] = mode.getGraphicsModeString()
	}
	g.Print(fmt.Sprintf(setGraphics, strings.Join(result, ";")))
}

func (g *Graphics) SetForegroundColor(color Color) {
	g.Print(fmt.Sprintf(setGraphics, setForegroundColor.addColorOffset(color)))
}

func (g *Graphics) SetBackgroundColor(color Color) {
	g.Print(fmt.Sprintf(setGraphics, setBackgroundColor.addColorOffset(color)))
}

func (g *Graphics) ResetForegroundColor() {
	g.Print(fmt.Sprintf(setGraphics, ResetForegroundColor.getGraphicsModeString()))
}

func (g *Graphics) ResetBackgroundColor() {
	g.Print(fmt.Sprintf(setGraphics, ResetBackgroundColor.getGraphicsModeString()))
}

func (g *Graphics) PrintAt(x, y int, s string, restorePosition bool) {
	if restorePosition {
		g.SavePosition()
		defer g.RestorePosition()
	}
	g.SetCursor(x, y)
	g.Print(s)
}

func (g *Graphics) DrawRectangle(x, y, width, height int, symbol rune) error {
	g.SavePosition()
	defer g.RestorePosition()
	termW, termH := terminal.GetTerminalSize()
	if x < 0 || x > termH {
		return fmt.Errorf("wrong x coordinate")
	}
	if y < 0 || y > termW {
		return fmt.Errorf("wrong Y coordinate")
	}

	for i := y; i < y+width; i++ {
		g.PrintAt(x, i, string(symbol), false)
		g.PrintAt(x+height, i, string(symbol), false)
	}

	for i := x; i < x+height; i++ {
		g.PrintAt(i, y, string(symbol), false)
		g.PrintAt(i, y+width, string(symbol), false)
	}
	return nil
}
