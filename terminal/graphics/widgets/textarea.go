package widgets

//
import (
	"log"
	"strings"

	"github.com/KlyuchnikovV/cui"
	"github.com/KlyuchnikovV/cui/server"
	"github.com/KlyuchnikovV/cui/types"
	"github.com/KlyuchnikovV/lines_buffer"
	"github.com/KlyuchnikovV/termin/keys"
)

type Textarea struct {
	baseElement
	buffer lines_buffer.Buffer
}

// TODO: make getting relative cursor position more easy
// TODO: move cursor to current area after every update (save/restore)
// TODO: optimize subscription mechanism

func NewTextarea(c *cui.ConsoleUI) *Textarea {
	t := &Textarea{
		baseElement: *newBaseElement(c, nil),
		buffer:      *lines_buffer.NewBuffer(""),
	}
	c.SubscribeWidget(server.KeyboardChan, t)
	c.SubscribeWidget(server.ResizeChan, t)
	return t
}

func (t *Textarea) Render(msg types.Message) {
	msg.Exec(t)

	switch msg.(type) {
	case *types.KeyboardMsg:
		value, ok := t.GetOption("rune").(keys.KeyboardKey)
		if !ok {
			log.Println("WARN: rune not set")
			return
		}
		log.Printf("textarea got rune %v\n", value)
		switch typed := value.(type) {
		case keys.RuneKey:
			t.processRune(typed)
		case keys.EscapeSequence:
			t.processEscapeSequence(typed)
		}
	case *types.ResizeMsg:
		t.printLines()
	}
}

func (t *Textarea) getAreaPosition() (int, int) {
	x, y := t.GetIntOption("x"), t.GetIntOption("y")
	return x + 1, y + 1
}

func (t *Textarea) ClearScreen() {
	w, h := t.GetIntOption("w"), t.GetIntOption("h")

	t.SavePosition()
	var replaceString = strings.Repeat(" ", w-2)
	for i := 0; i < h-2; i++ {
		t.PrintAt(i, 0, replaceString, false)
	}
	t.RestorePosition()
}

func (t *Textarea) printLines() {
	for i, line := range t.buffer.Lines() {
		t.PrintAt(i, 0, line, false)
	}
}

func (t *Textarea) SetCursor(x, y int) {
	xAbs, yAbs := t.getAreaPosition()
	t.Server.SetCursor(xAbs+x, yAbs+y)
}

func (t *Textarea) PrintAt(x, y int, text string, restorePosition bool) {
	xAbs, yAbs := t.getAreaPosition()
	t.Server.PrintAt(xAbs+x, yAbs+y, text, restorePosition)
}

func (t *Textarea) processRune(r keys.RuneKey) {

	switch {
	case r == keys.EndOfTransmission:
		log.Printf("TEXTAREA: print lines:\n%s", t.buffer.String())
	case r == keys.LineFeed:
		t.buffer.NewLine()
	case r == keys.Delete:
		t.buffer.DeleteBackward()
	case keys.IsSymbolChar(r):
		t.buffer.Insert(string(r.Rune()))
	}

	t.ClearScreen()
	t.printLines()
	t.SetCursor(t.buffer.RowNum(), t.buffer.ColumnNum())
}

func (t *Textarea) processEscapeSequence(e keys.EscapeSequence) {
	switch e {
	case keys.UpArrow:
		t.buffer.PrevLine()
		t.SetCursor(t.buffer.RowNum(), t.buffer.ColumnNum())
	case keys.DownArrow:
		t.buffer.NextLine()
		t.SetCursor(t.buffer.RowNum(), t.buffer.ColumnNum())
	case keys.LeftArrow:
		t.buffer.PrevRune()
		t.SetCursor(t.buffer.RowNum(), t.buffer.ColumnNum())
	case keys.RightArrow:
		t.buffer.NextRune()
		t.SetCursor(t.buffer.RowNum(), t.buffer.ColumnNum())
	case keys.DeleteKey:
		t.buffer.DeleteForward()
	}
}
