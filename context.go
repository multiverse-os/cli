package cli

type Input struct {
	Command *Command
	Flags   []*Flag
}

func LoadInput(command *Command, flags []*Flag) *Input {
	return &Input{Command: command, Flags: flags}
}
