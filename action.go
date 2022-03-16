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

func (self *Context) Execute() *Context {
  // NOTE: Run each action
  for _, action := range self.Actions {
    action(self)
  }
  return self
}
