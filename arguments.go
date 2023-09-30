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

// /////////////////////////////////////////////////////////////////////////////
// TODO: Need to be able to convert string[] to arguments
type arguments []Argument

func Arguments(arguments ...Argument) (argumentPointers arguments) {
	for index, _ := range arguments {
		argumentPointers = append(argumentPointers, arguments[index])
	}
	return argumentPointers
}

func (self arguments) Last() Argument  { return self[self.Count()-1] }
func (self arguments) First() Argument { return self[0] }
func (self arguments) Count() int      { return len(self) }

func (self arguments) HasNext(index int) bool {
	// NOTE
	// Index stars from zero len() doesn't, so we check
	// index + 2 to compensate for it starting at 0 and
	// jumping ahead
	return (index + 2) == len(self)
}

func (args arguments) Add(newArgument Argument) arguments {
	return append(append(arguments{}, newArgument), args...)
}
