package widgets

import (
	"log"

	"github.com/KlyuchnikovV/cui"
	"github.com/KlyuchnikovV/cui/types"
)

type Div struct {
	baseElement
	text string
}

func NewDiv(c *cui.ConsoleUI, text string) *Div {
	return &Div{
		baseElement: *newBaseElement(c, nil),
		text:        text,
	}
}

func (d *Div) Render(msg types.Message) {
	msg.Exec(d)
	x, y, w, h := d.GetIntOption("x"), d.GetIntOption("y"), d.GetIntOption("w"), d.GetIntOption("h")
	log.Printf("%s draw at %d %d %d %d\n", d.text, x, y, h, w)
	d.PrintAt(x+h/2, y+w/2, d.text, true)
}
