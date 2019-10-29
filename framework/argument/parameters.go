package argument

import (
	data "github.com/multiverse-os/cli/framework/data"
)

type Params struct {
	Chain    *Chain
	Position int
	Type     data.Type
	Value    []string
	//FileExt string // This is for autocomplete
}

func (self Params) PreviousCommand() (Command, bool) {
	return self.Chain.PreviousCommand(self.Position)
}

func (self Params) NextArguments() (Argument, bool) {
	return self.Chain.NextArgument(self.Position)
}
