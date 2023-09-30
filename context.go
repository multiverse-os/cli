package cli

type Context struct {
	CLI       *CLI
	Process   process
	Command   *Command
	Arguments arguments
	Commands  commands
	Flags     flags
	Params    params
	Actions   actions
}

func (ctx Context) Flag(name string) *Flag {
	return ctx.Flags.Name(name)
}

func (ctx Context) HasFlag(name string) bool {
	return ctx.Flag(name) != nil
}

// TODO
// This is the object we are giving the developer it should at least
// have matching Command and HasCommand, but likely a lot more.
// Like something to get the root psuedo-command (app)
