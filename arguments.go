package cli

type ArgumentType int

const (
  CommandArgument ArgumentType = iota
  FlagArgument
  ParamArgument
)

type Argument interface {
  Type()    ArgumentType
  IsValid() bool
}

type arguments []*Argument 

func Arguments(arguments ...Argument) (argumentPointers arguments) { 
  for index, _ := range arguments {
    argumentPointers = append(argumentPointers, &arguments[index])
  }
  return argumentPointers
}

func (self arguments) Last() *Argument { return self[self.Count()-1] }
func (self arguments) Count() int { return len(self) }

func (self arguments) Add(argument Argument) arguments {
  return append(self, &argument)
}
