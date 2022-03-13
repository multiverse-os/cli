package cli

type Argument interface {
  IsValid() bool
}

func ToParam(param Argument) *Param { return param.(*Param) }

func ToParams(paramArguments []Argument) (newParams params) {
  for _, paramArgument := range paramArguments {
    newParams = append(newParams, ToParam(paramArgument))
  }
  return newParams
}

func ToFlag(flag Argument) *Flag { return flag.(*Flag) }

func ToFlags(flagArguments []Argument) (newFlags flags) {
  for _, flagArgument := range flagArguments {
    newFlags = append(newFlags, ToFlag(flagArgument))
  }
  return newFlags
}

func ToCommand(command Argument) *Command { return command.(*Command) }

func ToCommands(commandArguments []Argument) (newCommands commands) {
  for _, commandArgument := range commandArguments {
    newCommands = append(newCommands, ToCommand(commandArgument))
  }
  return newCommands
}

func Reverse(arguments []Argument) (reversedArguments []Argument) {
  // TODO: Convert all for loops to this declaration format, its MUCH better
  // than the traditional one that exists for backwards compatibility
  for reversedIndex := len(arguments) - 1; reversedIndex >= 0; reversedIndex-- {
    reversedArguments = append(reversedArguments, arguments[reversedIndex])
  }
  return reversedArguments
}

///////////////////////////////////////////////////////////////////////////////
type arguments []Argument 

func Arguments(arguments ...Argument) (argumentPointers arguments) { 
  for index, _ := range arguments {
    argumentPointers = append(argumentPointers, arguments[index])
  }
  return argumentPointers
}

func (self arguments) Last() Argument { return self[self.Count()-1] }
func (self arguments) First() Argument { return self[0] }
func (self arguments) Count() int { return len(self) }

func (self arguments) Add(newArgument Argument) (prepended arguments) { 
  //prepended := Arguments(newArgument)
  return append(append(prepended, newArgument), self...)
}

func (self arguments) PreviousIfFlag() *Flag {
  argument := self.First()
	switch argument.(type) {
	case *Flag:
    return ToFlag(argument)
  default:
    return nil
  }
}
