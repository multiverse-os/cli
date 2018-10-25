package log

type Action func()
type HookType int

const (
	BeforeAction HookType = iota
	AfterAction
)

type Hook struct {
	Levels []LogLevel
	Type   HookType

	Action
}

func (self *Logger) AddHook(levels []LogLevel, hookType HookType, action Action) {
	hook := &Hook{
		Levels: levels,
		Type:   hookType,
		Action: action,
	}
	for _, level := range levels {
		self.Hooks[level][hookType] = append(self.Hooks[level][hookType], hook)
	}
}

func (self *Logger) ClearHooks() {
	self.Hooks = map[LogLevel]map[HookType][]*Hook{}
}

func (self *Logger) ExecuteHooks(level LogLevel, hookType HookType) {
	for _, hook := range self.Hooks[level][hookType] {
		hook.Action()
	}
}
