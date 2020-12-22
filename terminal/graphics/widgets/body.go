package widgets

import (
	"log"
	"os"
	"syscall"

	"github.com/KlyuchnikovV/cui"
	"github.com/KlyuchnikovV/cui/graphics"
	"github.com/KlyuchnikovV/cui/low_level/terminal"
	"github.com/KlyuchnikovV/cui/server"
	"github.com/KlyuchnikovV/cui/types"
)

type Body struct {
	baseElement
}

func NewBody(c *cui.ConsoleUI, children ...types.Widget) *Body {
	b := &Body{
		baseElement: *newBaseElement(c, nil, children...),
	}
	c.SubscribeWidget(server.ResizeChan, b)
	return b
}

func (b *Body) Render(msg types.Message) {
	b.SavePosition()
	msg.Exec(b)
	switch b.GetOption("signal").(os.Signal) {
	case syscall.SIGWINCH:
		b.ClearScreen(graphics.ClearAll)
		b.options["w"], b.options["h"] = terminal.GetTerminalSize()
		log.Printf("BODY: w: %d, h: %d", b.GetIntOption("w"), b.GetIntOption("h"))
		for _, child := range b.children {
			child.Render(types.NewResizeMsg(b.GetIntOption("x"), b.GetIntOption("y"), b.GetIntOption("w"), b.GetIntOption("h")))
		}
	}
	b.RestorePosition()
}
