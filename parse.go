package cli

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

func (self *CLI) Parse(arguments []string) (*Context, error) {
	defer self.benchmark(time.Now(), "benmarking argument parsing and action execution")
	cwd, executable := filepath.Split(arguments[0])

	context := &Context{
		CLI:          self,
		CWD:          cwd,
		Command:      &self.Command,
		Executable:   executable,
		CommandChain: &Chain{},
		Params:       Params{},
		Flags:        make(map[string]*Flag),
		Args:         arguments[1:],
	}
	context.CommandChain.AddCommand(&self.Command)

	for index, argument := range context.Args {
		if flagType, ok := HasFlagPrefix(argument); ok {
			fmt.Println("index: [%v] \n", index)
			fmt.Println("argument:", argument)

			// TODO: Need to handle skipping next argument when next argument is used
			context.ParseFlag(flagType, argument, context.NextArgument(index))

			//context.ParseFlag(index, flagType, &Flag{Name: argument})
		} else {
			if command, ok := context.Command.Subcommand(argument); ok {
				command.Parent = context.Command
				context.Command = &command
				context.CommandChain.AddCommand(context.Command)
			} else {
				for _, param := range context.Args[index:] {
					if flagType, ok := HasFlagPrefix(param); ok {
						context.ParseFlag(flagType, argument, context.NextArgument(index))
					} else {
						context.Params.Value = append(context.Params.Value, param)
					}
				}
				break
			}
		}
	}

	if context.CommandChain.UnselectedCommand() {
		context.Command = &Command{
			Parent: context.Command,
			Name:   "help",
		}
	}

	self.Debug = context.HasFlag("debug")
	if context.Command.is("version") || context.HasGlobalFlag("version") {
		self.RenderVersionTemplate()
	} else if context.Command.is("help") || context.HasFlag("help") {
		context.RenderHelpTemplate()
	} else {
		context.Command.Action(context)
	}
	return context, nil
}

//func (self *Context) ParseFlag(index int, flagType FlagType, flag *Flag) {
//	var flagParts []string
//	flagParts = strings.Split(flag.Name, Equal.String())
//	if 1 < len(flagParts) {
//		fmt.Println("there is more than 1 flagpart")
//		flag.Value = flagParts[1]
//	} else {
//		if len(self.Args) > index+1 {
//			fmt.Println("there is more than 1 flagpart")
//			flag.Value = self.Args[index+1]
//		} else {
//			fmt.Println("no flag part, assuming a value of 1")
//			flag.Value = "1"
//			flag.Type = data.Bool
//		}
//	}
//
//	fmt.Println("index is:", index)
//	fmt.Println("flagType:", flagType.String())
//	fmt.Println("flag:", flag)
//
//	if flagType == Short {
//		fmt.Println("flagType == Short:")
//
//		shortName := flagParts[0][1:]
//		// Stacked Flags
//		// TODO: Needs to work from specific to global so may need a for loop
//		// with minus index i--
//		for index, stackedFlag := range shortName {
//			// Load flag
//			if flagDefinition := self.Flag(string(stackedFlag)); flagDefinition != nil {
//				flag.Name = flagDefinition.Name
//				if index != len(flag.Name)-1 {
//					// NOTE: Stacked flag that is not the last element MUST be bool
//					flag.Value = "1"
//					flag.Type = data.Bool
//				} else {
//					// NOTE: Stacked flag that is last element needs to use value
//
//				}
//				self.ParseFlag(index, flagType, flag)
//			}
//		}
//	} else if flagType == Long {
//		flag.Name = flagParts[0][2:]
//	}
//}

func (self *Context) ParseFlag(flagType FlagType, argument, nextArgument string) (parsedFlag Flag) {
	argument = strings.ToLower(argument)
	// NOTE: Next argument may be value for flag so it may not be lowercased by
	//       default.
	//nextArgument = strings.ToLower(argument) -

	var flagParts []string
	if flagType == Short {
		// NOTE: Subtract dashes
		flagParts = strings.Split(argument[1:len(argument)], "=")

		// TODO: Handle stacking short flag
		if 1 <= len(flagParts[0]) {
			// NOTE: If a short tag is longer than 1 character
			for index, stackedFlag := range flagParts[0] {
				fmt.Println("stackedFlag:", stackedFlag)
				if index == len(flagParts[0]) {

				}
			}
			// TODO: Last item in stacked flags could be not boolean, check next
			// argumetn to decide
		}

	} else if flagType == Long {
		flagParts = strings.Split(argument[2:len(argument)], "=")
	}
	if len(flagParts) == 2 {
		parsedFlag.Name = flagParts[0]
		parsedFlag.Value = flagParts[1]
	} else {
		parsedFlag.Name = flagParts[0]
		if _, ok := HasFlagPrefix(nextArgument); ok {
			// NOTE: Next argument is flag, so flag.Value is 1 (boolean)
			parsedFlag.Value = "1"
		}
		for _, subcommand := range self.Command.Subcommands {
			if nextArgument == subcommand.Name {
				// NOTE: Next argument is a command, so flag.Value is 1 (boolean)
				parsedFlag.Value = "1"
			}
		}
	}
	if len(parsedFlag.Value) == 0 {
		fmt.Println("assigning nextArgument to parsedFlag.Value:", nextArgument)
		parsedFlag.Value = nextArgument
	}
	// TODO: This fucking sucks, everytime we update a flag we rebuild all the
	// commands flags.
	for _, command := range self.CommandChain.Reversed() {
		var flags []Flag
		for _, flag := range command.Flags {
			if flag.is(parsedFlag.Name) {
				fmt.Println("assinging parsedFlag to flag.Value:", parsedFlag.Value)
				flag.Value = parsedFlag.Value
			}
			flags = append(flags, flag)

		}
		command.Flags = flags

	}
	return parsedFlag
}

func FlagNameForType(flagType FlagType, argument string) (name string) {
	switch flagType {
	case Short:
		name = argument[1:len(argument)]
	case Long:
		name = argument[2:len(argument)]
	}
	return strings.ToLower(strings.Split(name, "=")[0])
}

func (self *Context) NextArgument(index int) string {
	if index+2 < len(self.Args) {
		return self.Args[index+1]
	}
	return ""
}
