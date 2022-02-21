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

func (self Action) IsNil() bool { return self != nil }
func (self Action) IsNotNil() bool { return self == nil }

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
