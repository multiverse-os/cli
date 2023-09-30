package cli

// TODO
// If we return context and error then onStart changes to context will be fixed
// but then we would need some sort of way of storing data between them like
// maybe a map
// TODO
// Also need a way to indicate that an action is final, or if it should be
// allowed to do the action of the parent command
type Action func(context *Context) error

type Actions struct {
	OnStart  Action
	Fallback Action
	OnExit   Action
}

type actions []Action

func (self *actions) Add(action Action) {
	if action != nil {
		*self = append(*self, action)
	}
}

//func (self actions) HasAction(action Action) bool {
//  for _, definedAction := range self {
//    if reflect.DeepEqual(definedAction, action) {
//      return true
//    }
//  }
//  return false
//}
