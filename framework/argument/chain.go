package argument

import (
	"fmt"
)

type Chain struct {
	Arguments []Argument
}

func ParseChain(arguments []string) *Chain {
	argumentChain := []Argument{}
	for _, argument := range arguments {
		argumentChain = append(argumentChain, ParseArgument(argument))
	}
	return &Chain{
		Arguments: argumentChain,
	}
}

// Validations ////////////////////////////////////////////////////////////////
func (self *Chain) IsValidPosition(position int) bool {
	return (1 < position && position <= len(self.Arguments))
}

// Routing ////////////////////////////////////////////////////////////////////
func (self *Chain) CommandPath() (path []string) {
	for _, argument := range self.Arguments {
		switch argument.(type) {
		case Command:
			path = append(path, argument.(Command).Arg)
		}
	}
	return path
}

// Argument Selection & Filtering //////////////////////////////////////////////
func (self *Chain) NextArgument(position int) (Argument, bool) {
	if self.IsValidPosition(position) {
		return self.Arguments[position+1], true
	}
	return nil, false
}

func (self *Chain) PreviousCommand(position int) (Command, bool) {
	if self.IsValidPosition(position) {
		for i := (len(self.Arguments) - 1); i >= 0; i-- {
			argument := self.Arguments[i]
			switch argument.(type) {
			case Command:
				if argument.(Command).Position < position {
					return argument.(Command), true
				}
			}
		}
	}
	return Command{}, false
}

func (self *Chain) TrailingFlags(position int) Flags {
	flagGroup := Flags{}
	if self.IsValidPosition(position) {
		for i := (len(self.Arguments) - 1); i >= position; i-- {
			element := self.Arguments[i]
			switch element.(type) {
			case Flag:
				fmt.Println("[framework/argument] adding flag to group of trailing flags:", element.(Flag).Value)
				flagGroup.Add(element.(Flag))
			}
		}
	}
	return flagGroup
}
