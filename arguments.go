package cli

type Argument interface {
  IsValid() bool
}

func ToParam(param Argument) *Param { return param.(*Param) }
func ToFlag(flag Argument) *Flag { return flag.(*Flag) }
func ToCommand(command Argument) *Command { return command.(*Command) }

///////////////////////////////////////////////////////////////////////////////
type arguments []Argument 

func Arguments(arguments ...Argument) (argumentPointers arguments) { 
  for index, _ := range arguments {
    argumentPointers = append(argumentPointers, arguments[index])
  }
  return argumentPointers
}

func (self arguments) Last() Argument { return self[self.Count()-1] }
func (self arguments) Count() int { return len(self) }

func (self arguments) Add(argument Argument) arguments { 
  return append(self, argument)
}

func (self arguments) Reversed() (reversedArguments arguments) {
  for i := self.Count() - 1; i >= 0; i-- {
    reversedArguments = append(reversedArguments, self[i])
  }
  return reversedArguments
}

// TODO: This works but we would rather build the prepend function, get rid of
// Reversed() if we don't end up using it, 
func (self arguments) PreviousFlag() *Flag {
  argument := self.Reversed()[0]
	switch argument.(type) {
	case *Flag:
    return ToFlag(argument)
  default:
    return nil
  }
}
