package cli

type Context struct {
	PID          int
	CLI          *CLI
	CWD          string
	Executable   string
	Command      *Command
	Params       Params
	CommandChain *Chain
	Flags        map[string]*Flag
	Args         []string
}

func (self *Context) HasCommandFlag(name string) bool {
	return self.CommandFlag(self.Command.Name, name) != nil
}

func (self *Context) HasFlag(name string) bool       { return self.Flag(name) != nil }
func (self *Context) HasGlobalFlag(name string) bool { return self.GlobalFlag(name) != nil }
func (self *Context) HasGlobalFlags() bool           { return 0 < len(self.GlobalFlags()) }
func (self *Context) HasCommands() bool              { return self.CommandChain.IsRoot() && self.HasSubcommands() }
func (self *Context) HasSubcommands() bool           { return 0 < len(self.Command.Subcommands) }

func (self *Context) HasSubcommand(name string) bool {
	_, hasSubcommand := self.Subcommand(name)
	return hasSubcommand
}

func (self *Context) Subcommand(name string) (*Command, bool) {
	for _, subcommand := range self.Command.Subcommands {
		if subcommand.is(name) {
			return &subcommand, true
		}
	}
	return nil, false
}

func (self *Context) GlobalFlag(name string) (flag *Flag) {
	for _, flag := range self.GlobalFlags() {
		if flag.is(name) {
			if len(flag.Value) == 0 {
				flag.Value = flag.Default
			}
			return flag
		}
	}
	return flag
}

func (self *Context) Flag(name string) (flag *Flag) {
	if self.HasCommands() {
		for _, command := range self.CommandChain.Reversed() {
			for _, flag := range command.Flags {
				if flag.is(name) {
					if len(flag.Value) == 0 {
						flag.Value = flag.Default
					}
					return &flag
				}
			}
		}
	}
	return flag
}

func (self *Context) CommandFlag(commandName, flagName string) (flag *Flag) {
	for _, command := range self.CommandChain.Commands {
		if command.is(commandName) {
			for _, flag := range command.Flags {
				if flag.is(flagName) {
					if len(flag.Value) == 0 {
						flag.Value = flag.Default
					}
					return &flag
				}
			}
		}
	}
	return flag

}

func (self *Context) GlobalFlags() map[string]*Flag {
	flags := make(map[string]*Flag)
	for _, flag := range self.CommandChain.First().Flags {
		flags[flag.Name] = &flag
	}
	return flags
}

func (self *Context) UpdateFlag(name, value string) {
	if self.HasCommands() {
		for _, command := range self.CommandChain.Reversed() {
			for _, flag := range command.Flags {
				if flag.Name == name || flag.Alias == name {
					flag.Value = value
				}
			}
		}
	}
}
