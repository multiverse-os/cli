package widgets

import (
	"github.com/KlyuchnikovV/cui"
	"github.com/KlyuchnikovV/cui/types"
)

type baseElement struct {
	*cui.ConsoleUI
	options  map[string]interface{}
	children []types.Widget
}

func newBaseElement(c *cui.ConsoleUI, options map[string]interface{}, children ...types.Widget) *baseElement {
	if options == nil {
		options = make(map[string]interface{})
	}
	return &baseElement{
		ConsoleUI: c,
		options:   options,
		children:  children,
	}
}

func (b *baseElement) SetOptions(opts map[string]interface{}) {
	for key, opt := range opts {
		b.options[key] = opt
	}
}

func (b *baseElement) GetOption(s string) interface{} {
	return b.options[s]
}

func (b *baseElement) GetOptions() map[string]interface{} {
	return b.options
}

func (b *baseElement) GetIntOption(s string) int {
	opt := b.options[s]
	if opt == nil {
		return 0
	}
	result, ok := opt.(int)
	if !ok {
		return 0
	}
	return result
}
