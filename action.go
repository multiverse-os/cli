package cli

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
