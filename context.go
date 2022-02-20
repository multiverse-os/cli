package cli

// NOTE: The reason hooks & actions exists within the context is that these are
// the specifed actions and hooks being ran (specified by the args passed from
// the os when the command was ran). CLI contains all possible hooks and
// actions, and context contains the actions and hooks being executed. 
type Context struct {
	PID          int
	CLI          *CLI
	CWD          string
	Executable   string
	Command      *Command
	Params       Params
	CommandChain *Chain
  Hooks        Hooks
  Actions      actions
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
  // TODO: DO we get a preformance increase and how much from hardcoding in
  // byte(\0x or whatever it is and doing the bytewise comparison? 
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
				return flag
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
					return flag
				}
			}
		}
	}
	return flag
}

// TODO: I would prefer API to be 

//     self.Command(1), or at least self.Command.First(),... CommandChain is
//     descriptive but clunky
//func (self *Context) HasCommand(name string) bool {
//	_, hasCommand := self.Command(name)
//	return hasCommand
//}
//
//func (self *Context) Command(name string) (*Command, bool) {
//	for _, subcommand := range self.Command.Subcommands {
//		if subcommand.is(name) {
//			return &subcommand, true
//		}
//	}
//	return nil, false
//}

// TODO: This is nice to have in its own function because this piece dictates
// core aspect of the logic and this puts it in a capsule easily understood or
// changed hopefully. 

// TODO: Global and Fallback actions are both running, this should be impossible
//       by definition
func (self *Context) ExecuteActions() {

}

// TODO: Context should only logically hold the meta methods, nothing directly
// acting on any collection or object inside but more the helpers made from
// merging those lower level 

func (self *Context) HasGlobalAction() bool { return self.CLI.Actions.Global.IsNil() }
func (self *Context) HasFallbackAction() bool { return self.CLI.Actions.Fallback.IsNil() }

// TODO: This needs to be updated to reflect the changes made by splitting up
// hooks from generic actions
// TODO: Does not take into consideration the command chain and the parent
// commands potential actions or hooks! or even the current commands hooks, or
// the global hooks. Right now its just the global actions (fallback or global)
// and the last command in the commandchain's action (so not complete) 
func (self *Context) HasAction() bool {
  return self.HasGlobalAction()   || 
         self.HasFallbackAction() || 
         self.Command.HasAction() // TODO: This needs to iterate through
                                  //       each command in chain to determine IF
                                  //       it has a possible action to preform
                                  //       (will be stored in a similar to
                                  //       CLI.Actions,.. Command.Actions but it
                                  //       should have specific actions more
                                  //       meaningful to command, so they cant
                                  //       share the same object. 
                                }
