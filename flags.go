package cli

type Flag struct {
	//Command       *Command
	Name          string
	Alias         string
	Type          DataType
	Description   string
	Hidden        bool
	FileExtension string // Would help us autogenerate nice autocompletion
	DefaultValue  interface{}
}

// TODO: Could probably speed up lookup and avoid this by putting flag in a
// lookup map twice, once with name and once with alias and just use a symbol
func (self Flag) is(name string) bool { return self.Name == name || self.Alias == name }
func (self Flag) visible() bool       { return !self.Hidden }

func (self Flag) usage() (output string) {
	output += "--" + self.Name
	if len(self.Alias) > 0 {
		output += ", -" + self.Alias
	}
	return output
}

// Public Methods ////
func Flags(flags ...Flag) []Flag { return flags }
