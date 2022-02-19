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

func (self *Context) UpdateFlag(name, value string) {
	for _, command := range self.CommandChain.Reversed() {
		for _, flag := range command.Flags {
			if flag.is(name) {
				flag.Value = value
			}
		}
	}
}

func (self *Context) HasFlag(name string) bool {
	return self.Flag(name).Value != "0"
}

func (self *Context) HasCommandFlag(name string) bool {
	return self.CommandFlag(self.Command.Name, name) != nil
}

func (self *Context) Flag(name string) *Flag {
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

	return &Flag{
		Name:  name,
		Value: "0",
	}
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

// TODO: I would prefer API to be 

//     self.Command(1), or at least self.Command.First(),... CommandChain is
//     descriptive but clunky
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

// TODO: This is nice to have in its own function because this piece dictates
// core aspect of the logic and this puts it in a capsule easily understood or
// changed hopefully. 
func (self *Context) Action() {
  if self.Command.HasAction() {
    self.Command.Action(self)
  }else{
    if self.HasGlobalAction() {
      self.CLI.Actions.Global(self)
      // TODO: May need else to either add ehlp command or print help directly.
      // This could go before version maybe but then version would need and help
      // would probably have to have their actions moved to their
      // initializiation as default hidden options. And i ahvent decided if that
      // is better
    }
  }

  self.CLI.Actions.OnStart(self)
  if {

    // Command Action 
    // Subcommand action ... for each command up the chain (remember to use the
    // flags for: (each level)+(global) flags. 

    // Global

    
  }else{
    self.CLI.Actions.Fallback(self)
  }
  self.CLI.Actions.OnExit(self)
}

func (self *Context) HasGlobalAction() bool { return self.CLI.Actions.Global != nil }
func (self *Context) HasFallbackAction() bool { return self.CLI.Actions.Fallback != nil }
func (self *Context) HasOnStartAction() bool { return self.CLI.Actions.OnStart != nil }
func (self *Context) HasOnExitAction() bool { return self.CLI.Actions.OnExit != nil }

func (self *Context) HasAction() bool {
  return self.HasGlobalAction()   || 
         self.HasFallbackAction() || 
         self.HasOnStartAction()  || 
         self.HasOnExitAction()   ||
         self.Command.HasAction() // TODO: This needs to iterate through
                                  //       each command in chain to determine IF
                                  //       it has a possible action to preform
                                  //       (will be stored in a similar to
                                  //       CLI.Actions,.. Command.Actions but it
                                  //       should have specific actions more
                                  //       meaningful to command, so they cant
                                  //       share the same object. 

// TODO: Move this to command? Command.HasNoAction() ?
func (self *Context) HasNoAction() bool { !return self.HasNoDefaultAction() }
func (self *Context) HasNoCommands() bool {  return !self.Command.HasCommands() }
