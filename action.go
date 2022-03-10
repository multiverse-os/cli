package cli

// TODO: Right now I feel like should probably only have global hooks, because
// command hooks may be redudant or introduce a lot of complexity given they
// would need to hit all hooks of parent commands. 
type Action func(context *Context) error

type actions []*Action

// TODO: Add the merging of several non-pointed actions into the actions type
// which is a slice of action pointers (see flags or commands on how this is
// done)

func (self actions) Count() int { return len(self) }

//func (self Action) IsNil() bool { return self != nil }
//func (self Action) IsNotNil() bool { return self == nil }

// TODO: I would like it to output if it successfully ran, and the error
//func (self Action) Execute() error {

// TODO: Execute will probably work by populating the actions chain and then
// just processing it in order. 
//func (self Action) Execute(context *Context) {
//  if self.IsNotNil() {
//    if context.Hooks.BeforeAction.IsNotNil() {
//      context.Hooks.BeforeAction(context)
//    }
//      
//    // TODO: I believe this should work but Im not 100% sure on this
//    self(context)
//
//    if context.Hooks.AfterAction.IsNotNil() {
//      context.Hooks.AfterAction(context)
//    }
//  }
//}

// TODO: Add a execute function, that checks if its nil before executing, and
// perhaps also runs the hooks all in one. 
type Hooks struct {
  BeforeAction Action
  AfterAction  Action
}

type Actions struct {
  Global    Action
  Fallback  Action
  // OnExit? Or Close? or this just covered by After?
}

func (self *CLI) Execute() *CLI {
  // TODO: This is currently the router, it would be nice to be able to produce
  // a standard URL like output (even have a URI scheme, like 

  //  cli://user@program:/command/subcommand?params
  //  
  //  OR somethjing similar, then be able to route to a defined functions in a
  //  controller section, but additionally and importantly, provide consistent,
  //  specific and useful details to the controller function so that they can be
  //  slim and written similarly. 
  // 

  // TODO: Iterate over context.Actions and execute each action, because this
  // slice will be popualted during parse, and this new logic will never have
  // the issue of trying to run a struct field that is of type *Action and have
  // it be nil 


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
