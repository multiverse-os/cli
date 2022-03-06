package cli

type Argument interface {
  IsValid() bool
}

func ToParam(param Argument) Param { return param.(Param) }
func ToFlag(flag Argument) Flag { return flag.(Flag) }
func ToCommand(command Argument) Command { return command.(Command) }

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
