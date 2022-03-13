package cli

type Action func(context *Context) error

type Actions struct {
  OnStart   Action
  Fallback  Action
  OnExit    Action
}

type actions []Action

// NOTE: The other Add() functions prepend, whereas this appends.
func (self actions) Add(action Action) actions {
  return append(self, action)
}

func (self *CLI) Execute() *CLI {
  // TODO: This is currently the router, it would be nice to be able to produce
  // a standard URL like output (even have a URI scheme, like 

  //  cli://user@program:/command/subcommand?params

  //if context.Command.is("version") || context.HasFlag("version") {
	//	self.RenderVersionTemplate()
  //} else if context.HasFlag("help") { // TODO: Removed condition where subcommands but no action that should get help output BUT -- should default action run regardless or above happens only when no default
	//	  context.RenderHelpTemplate(context.Command)
  //} else if context.Command.is("help") {
	//	  context.RenderHelpTemplate(context.Command.Parent)
  //} else {
  //    // Produce a list of actions that need to be run and put them in the
  //    // context Chain object for later execution so it will eventually be
  //    // cli.Parse(os.Args).Execute() 
  //    //context.ExecuteActions()
	//}
  return self
}
