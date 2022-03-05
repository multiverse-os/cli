package cli

type Context struct {
	CLI          *CLI
  Process      process
  Command      *Command
  Arguments    arguments
  Commands     commands
  Flags        flags
  Params       params
}

func (self Context) Flag(name string) *Flag { 
  return self.Flags.Name(name)
}

func (self Context) HasFlag(name string) bool { 
  return self.Flag(name) != nil 
}
