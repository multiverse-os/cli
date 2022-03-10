package cli

type Argument interface {
  IsValid() bool
}



func ToParam(param Argument) *Param { return param.(*Param) }

func ToParams(paramArguments []Argument) (params params) {
  for _, paramArgument := range paramArguments {
    params = append(params, ToParam(paramArgument))
  }
  return params
}

func ToFlag(flag Argument) *Flag { return flag.(*Flag) }

func ToFlags(flagArguments []Argument) (flags flags) {
  for _, flagArgument := range flagArguments {
    flags = append(flags, ToFlag(flagArgument))
  }
  return flags
}

func ToCommand(command Argument) *Command { return command.(*Command) }

func ToCommands(commandArguments []Argument) (commands commands) {
  for _, commandArgument := range commandArguments {
    commands = append(commands, ToCommand(commandArgument))
  }
  return commands
}

func Reverse(arguments []Argument) (reversedArguments []Argument) {
  for index := len(arguments) - 1; index >= 0; index-- {
    reversedArguments = append(reversedArguments, arguments[index])
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

func (self arguments) PreviousFlag() *Flag {
  argument := self.First()
	switch argument.(type) {
	case *Flag:
    return ToFlag(argument)
  default:
    return nil
  }
}
