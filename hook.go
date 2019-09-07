package cli

type HookType int

const (
	BeforeAction HookType = iota
	AfterAction
)

type Hook struct {
	Type    HookType
	Command *Command
}
