package argument

type Argument interface {
	NextArgument() (Argument, bool)
}

func ParseArgument(argument string) Argument {
	if flagType, ok := IsValidFlag(argument); ok {
		return Flag{
			Identifier: flagType,
			Arg:        argument,
		}
	} else {
		// TODO: This is either command or parameter
		return Command{
			Arg: argument,
		}
	}
}
