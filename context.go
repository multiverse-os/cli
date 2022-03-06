package cli

import (
  "fmt"
)

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

// DEV
func (self *Context) DevOutput() {
  fmt.Println("================")
  fmt.Println("parsing COMPLETED!") 
  fmt.Println("arguments parsed: ", len(self.Arguments))
  fmt.Println("                  ", self.Arguments)
  fmt.Println("commands parsed:  ", len(self.Commands))
  fmt.Println("                  ", self.Commands)
  fmt.Println("flags parsed:     ", len(self.Flags))
  fmt.Println("                  ", self.Flags)
  fmt.Println("params parsed:    ", len(self.Params))
  fmt.Println("                  ", self.Params)
  fmt.Println("---------------")

  fmt.Println("Command(first)")
  fmt.Println("  Name:        ", self.Command.Name)
  fmt.Println("  Alias:       ", self.Command.Alias)
  fmt.Println("  Description: ", self.Command.Description)
  fmt.Println("  Hidden:      ", self.Command.Hidden)
  fmt.Println("  Parent:      ", self.Command.Parent)
  fmt.Println("  Subcommands: ", self.Command.Subcommands)
  fmt.Println("  Flags:       ", self.Command.Flags)
  fmt.Println("  Action:      ", self.Command.Action)
  fmt.Println("  Hooks:       ", self.Command.Hooks)
}

