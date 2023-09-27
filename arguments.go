package cli

import (
	"fmt"
)

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

func (args arguments) Add(newArgument Argument) arguments {
	//prepended := Arguments(newArgument)
	args = append(append(arguments{}, newArgument), args...)
	return args
}

// TODO: Are args here are wrong, we have ltierally nmber 1 and a pointer and
// then a full object, this is crazy shit

// AND fucking adding flags should be done as pointers again~!
func (args arguments) PreviousIfFlag() *Flag {
	fmt.Printf("arguments:(%v); len(%v)\n", args, len(args))
	for index, arg := range args {
		fmt.Printf("index(%v)=arg(%v)\n", index, arg)
		switch arg.(type) {
		case *Flag:
			fmt.Printf("is *Flag(%v)\n", arg)
		default:
			fmt.Printf("is nil\n")
		}
	}
	argument := args.First()
	switch argument.(type) {
	case *Flag:
		fmt.Printf("PREVIOUS ARGUMENT WAS A FLAG! arg(%v)\n", argument)
		return argument.(*Flag)
		//return ToFlag(argument)

	default:
		return nil
	}
}
