package log

import "fmt"

// TODO: This structure is wrong
type Hook interface {
	Levels() []LogLevel

	BeforeAction(*Entry) error
	AfterAction(*Entry) error
}

// TODO: It would be better to move this into Logger object instead of having a
// global variable accessible from anywhere.
type LevelHooks map[LogLevel][]Hook

// Add a hook to an instance of logger. This is called with
// `log.Hooks.Add(new(MyHook))` where `MyHook` implements the `Hook` interface.
func (hooks LevelHooks) Add(hook Hook) {
	for _, level := range hook.Levels() {
		hooks[level] = append(hooks[level], hook)
	}
}

func (hooks LevelHooks) BeforeAction(level LogLevel, entry *Entry) error {
	for _, hook := range hooks[level] {
		fmt.Println("hook: ", hook)
		// TODO: Execute Before Action Hook
		//if err := hook.Fire(entry); err != nil {
		//	return err
		//}
	}
	return nil
}

func (hooks LevelHooks) AfterAction(level LogLevel, entry *Entry) error {
	for _, hook := range hooks[level] {
		fmt.Println("hook: ", hook)
		//if err := hook.Fire(entry); err != nil {
		//	return err
		//}
	}
	return nil
}
