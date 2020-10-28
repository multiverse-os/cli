package token

import (
	"strings"
)

type Identifier int

const (
	Short Identifier = iota + 1
	Long
	NotAvailable
)

func (self Identifier) Is(identifier Identifier) bool { return self == identifier }
func (self Identifier) Length() int                   { return int(self) }
func (self Identifier) String() string                { return strings.Repeat("-", self.Length()) }

type Separator int

const (
	Whitespace Separator = iota
	Equal
)

func (self Separator) Is(separator Separator) bool { return self == separator }

func (self Separator) String() string {
	if self.Is(Equal) {
		return "="
	} else {
		return " "
	}
}
