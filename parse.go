package cli

type FlagType int

const (
	BoolFlag FlagType = iota
	IntFlag
	StringFlag
	PathFlag
	FilenameFlag
)

type Flag struct {
	Name    string // Primary name
	Aliases []string
	Type    FlagType
	Usage   string
	Hidden  bool
	Value   interface{}
}
type Flags []Flag

func (flag Flag) Names() []string { return append([]string{flag.Name}, flag.Aliases...) }

func flagPrefix(name string) string {
	if len(name) == 1 { // TODO: And two?
		return "-"
	} else {
		return "--"
	}
}

// TODO: Are hooks really necessary? Maybe it would be better to just implement
// a middleware like functionality and push this even closer to being more like
// web development to make it easier to comphrehend and extend
// TODO: Why do we have 'Usage' AND 'UsageText' seems like we should be merging this in some way. Also is this diff than description?
type Command struct {
	Hidden        bool
	Category      int
	Name          string
	Aliases       []string
	ParentCommand *Command
	Subcommands   map[string]Command
	Flags         map[string]Flag
	Usage         string
	Action        interface{}
}

type Commands []Command

func (self Command) HasSubcommands() bool { return len(self.Subcommands) > 0 }
func (self Command) Names() []string      { return append([]string{self.Name}, self.Aliases...) }
