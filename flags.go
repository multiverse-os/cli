package cli

const (
	longFlag       = "--"
	shortFlag      = "-"
	valueDelimeter = "=" // TODO: Need to support space too
)

type Flag struct {
	//Command       *Command
	Name        string
	Alias       string
	Description string
	Hidden      bool
	Type        DataType
	Value       interface{}
	//FileExt     string // Would help us autogenerate nice autocompletion
}

// TODO: Could probably speed up lookup and avoid this by putting flag in a
// lookup map twice, once with name and once with alias and just use a symbol
func (self Flag) is(name string) bool { return self.Name == name || self.Alias == name }
func (self Flag) visible() bool       { return !self.Hidden }

func (self Flag) usage() (output string) {
	output += longFlag + self.Name
	if len(self.Alias) > 0 {
		output += ", " + shortFlag + self.Alias
	}
	return output
}

//
// Flag Input
///////////////////////////////////////////////////////////////////////////////
type inputFlag struct {
	command *inputCommand
	Type    DataType
	Name    string
	Value   string
}

func newInputFlag() *inputFlag {
	return &inputFlag{
		Value: "1",
	}
}

// Flags //////////////////////////////////////////////////////////////////////
type inputFlags []*inputFlag

func newFlagGroup() *inputFlags { return &inputFlags{} }

func (self *inputFlags) add(flag *inputFlag) { *self = append(*self, flag) }
func (self *inputFlags) reset()              { self = &inputFlags{} }
func (self *inputFlags) isEmpty() bool       { return IsZero(len(*self)) }

//
// Public Methods
///////////////////////////////////////////////////////////////////////////////
func Flags(flags ...Flag) []Flag { return flags }
