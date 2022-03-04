package cli

// NOTE: The reason hooks & actions exists within the context is that these are
// the specifed actions and hooks being ran (specified by the args passed from
// the os when the command was ran). CLI contains all possible hooks and
// actions, and context contains the actions and hooks being executed. 
type Context struct {
	CLI          *CLI
  Process      process

  Arguments    arguments

  Command      *Command
  Commands     commands
  Flags        flags
  Params       params

  // TODO: It may be wise to make both Chain and Actions private, to guarantee
  // use via the cached versions above (and continue considering putting actions
  // into chain (action chain no?) 
	Chain       *Chain
  Debug       bool
}

// TODO: Add ability to access Banner/Spinner and other cli text user interface
// (TUI) tools

// TODO: Need a mirror function in CLI for pulling out defined flags

//     c.Flags["debug"].Bool() -> c.Flag("Debug").Bool()
func (self Context) Flag(name string) *Flag { 
  return self.Chain.Flags.Name(name)
}

//Params       params
// TODO: Perhaps change name to Arguments to create the API
//          c.Arguments.Flags() 
//          c.Arguments.Command() -> Last command in argument chain
//          c.Arguments.Actions() -> produces list of actions, removing it
//          from BOTH context and chain
//Hooks     Hooks
//Actions   actions
//Args      []string
// c.Flags c.Command c.Flags

// TODO: Move Reversed() in arguments to commands object as a method
// TODO: Im about to break down this logic and seperate it- but the original
// meta task will still need to be capable of being done and so im leaving this
// here to remind myself to rebuild this function with the new fucntions we
// create by breaking it apart
//func (self *Context) UpdateFlag(name, value string) {
//	for _, command := range self.Chain.Reversed() {
//		for _, flag := range command.Flags {
//			if flag.is(name) {
//				flag.Value = value
//			}
//		}
//	}
//}

func (self Context) HasFlag(name string) bool { 
  return self.Flag(name) != nil 
}

//func (self *Context) Flag(name string) *Flag {
//	for _, command := range self.CommandChain.Reversed() {
//		for _, flag := range command.Flags {
//			if flag.is(name) {
//				if len(flag.Value) == 0 {
//					flag.Value = flag.Default
//				}
//				return flag
//			}
//		}
//	}
//
//	return &Flag{
//		Name:  name,
//		Value: "0",
//	}
//}

// TODO: This logic should just be pushed onto Command.Flag("name")
//func (self *Context) CommandFlag(commandName, flagName string) (flag *Flag) {
//	for _, command := range self.CommandChain.Commands {
//		if command.is(commandName) {
//			for _, flag := range command.Flags {
//				if flag.is(flagName) {
//					if len(flag.Value) == 0 {
//						flag.Value = flag.Default
//					}
//					return flag
//				}
//			}
//		}
//	}
//	return flag
//}

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
//func (self *Context) ExecuteActions() {
//  // TODO: 
//  if self.CLI.Actions.Global.IsNotNil() {
//    self.Actions = append(self.Actions, self.CLI.Actions.Global)
//  }
//
//  //  iterate in over the commands in the command chain, and grab any
//  //  defined action and add it
//  //command -f command subcommand
//  //command command subcommand
//   
//
//
//  
//  //   check if we are running a command fucntion OR the fallback
//  //      ensure we are grabbing all the command functions if there is a chain
//  //      of commands 
//  //      ensure we are grabbing all the command functions hooks associated
//  //   grab all the global hooks and put them in the context.Hooks with the
//  //   command hooks so they are all there to be run together. 
//
//  //   Run before, global+(command actions|fallback), after
//  // 
//  //   Probably return an error 
//}

// TODO: Context should only logically hold the meta methods, nothing directly
// acting on any collection or object inside but more the helpers made from
// merging those lower level 

//func (self *Context) HasGlobalAction() bool { return self.CLI.Actions.Global.IsNil() }
//func (self *Context) HasFallbackAction() bool { return self.CLI.Actions.Fallback.IsNil() }

// TODO: This needs to be updated to reflect the changes made by splitting up
// hooks from generic actions
// TODO: Does not take into consideration the command chain and the parent
// commands potential actions or hooks! or even the current commands hooks, or
// the global hooks. Right now its just the global actions (fallback or global)
// and the last command in the commandchain's action (so not complete) 
//func (self *Context) HasAction() bool {
//  return self.HasGlobalAction()   || 
//         self.HasFallbackAction() || 
//         self.Command.HasAction() // TODO: This needs to iterate through
//                                  //       each command in chain to determine IF
//                                  //       it has a possible action to preform
//                                  //       (will be stored in a similar to
//                                  //       CLI.Actions,.. Command.Actions but it
//                                  //       should have specific actions more
//                                  //       meaningful to command, so they cant
//                                  //       share the same object. 
//                                }
