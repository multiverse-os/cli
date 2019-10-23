package cli

type Command struct {
	Hidden      bool
	Category    int
	Name        string
	Alias       string
	Parent      *Command
	Subcommands []Command
	Flags       []Flag
	Description string
	Action      func(c *Context) error
}

func (self Command) is(name string) bool  { return self.Name == name || self.Alias == name }
func (self Command) visible() bool        { return !self.Hidden }
func (self Command) isRootCommand() bool  { return self.Parent == nil }
func (self Command) hasSubcommands() bool { return len(self.Subcommands) == 0 }

func (self Command) usage() (output string) {
	if len(self.Alias) > 0 {
		output += ", " + self.Alias
	}
	return self.Name + output
}

func (self Command) visibleSubcommands() (commands []Command) {
	for _, command := range self.Subcommands {
		if command.visible() {
			commands = append(commands, command)
		}
	}
	return commands
}

// Public Methods ////
func Commands(commands ...Command) []Command { return commands }
