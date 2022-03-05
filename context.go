package cli

type Context struct {
	CLI          *CLI
  Process      process
  // Command, Flag, Param & Action Chain
	chain       chain
  // Cached Values
  Arguments    arguments
  Command      *Command
  Commands     commands
  Flags        flags
  Params       params
}

// TODO: Decide if we interact with flags via cached object or via function like
// below. The difference to the API would be:
//   c.Flags().Name("debug")
//   c.Flags.Name("debug")
//   c.Flag("debug")
//func (self Context) Flags() flags {
//  return self.chain.Flags
//}

func (self Context) Flag(name string) *Flag { 
  return self.chain.Flags.Name(name)
}

func (self Context) HasFlag(name string) bool { 
  return self.Flag(name) != nil 
}
