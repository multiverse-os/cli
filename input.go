package cli

type Input struct {
	CLI     *CLI
	Command *Command
	Flags   []*Flag
}

func LoadInput(cli *CLI, command *Command, flags []*Flag) *Input {
	return &Input{CLI: cli, Command: command, Flags: flags}
}
