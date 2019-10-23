package cli

// TODO: This is the object we are giving to the resulting action defined in the
// command.
type InputCommand struct {
	// TODO: Value could store the parameter for command, and value for flag
	Name          string
	ParameterType DataType
	Parameters    []interface{}
	Flags         map[string]*InputFlag
	Definition    *Command
	ParentCommand *InputCommand
	Subcommands   []*InputCommand
	depth         int
}

type InputFlag struct {
	Name    string
	Type    DataType
	Command *InputCommand
	Value   interface{}
}

// TODO: Maybe should consider moving the flag validation and datatype based
// output to Argument, and see if we can return only arguments to the action
// that is run as a result of the CLI execution
