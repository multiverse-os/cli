package cli

//"sync/atomic"

type CommandCategory struct {
	Name        string
	Description string
	Commands    Commands
	Hidden      bool
}
