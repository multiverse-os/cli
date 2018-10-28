package log

type Action func()
type HookType int

const (
	BeforeAction HookType = iota
	AfterAction
)

type Hook struct {
	levels   []LogLevel
	hookType HookType
	action   Action
}

func (self *Logger) AddHook(levels []LogLevel, hookType HookType, action Action) {
	hook := &Hook{
		levels:   levels,
		hookType: hookType,
		action:   action,
	}
	for _, level := range levels {
		self.hooks[level][hookType] = append(self.hooks[level][hookType], hook)
	}
}

func (self *Logger) ClearHooks() {
	self.hooks = map[LogLevel]map[HookType][]*Hook{}
}

func (self *Logger) ExecuteHooks(level LogLevel, hookType HookType) {
	for _, hook := range self.hooks[level][hookType] {
		hook.action()
	}
}
