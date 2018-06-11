package cli

// CommandCategories is a slice of *CommandCategory.
type CommandCategories []*CommandCategory

// CommandCategory is a category containing commands.
type CommandCategory struct {
	Name     string
	Commands Commands
}

func (self CommandCategories) Less(i, j int) bool {
	return lexicographicLess(self[i].Name, self[j].Name)
}

func (self CommandCategories) Len() int {
	return len(self)
}

func (self CommandCategories) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

// AddCommand adds a command to a category.
func (self CommandCategories) AddCommand(category string, command Command) CommandCategories {
	for _, commandCategory := range self {
		if commandCategory.Name == category {
			commandCategory.Commands = append(commandCategory.Commands, command)
			return self
		}
	}
	return append(self, &CommandCategory{Name: category, Commands: []Command{command}})
}

// VisibleCommands returns a slice of the Commands with Hidden=false
func (self *CommandCategory) VisibleCommands() []Command {
	ret := []Command{}
	for _, command := range self.Commands {
		if !command.Hidden {
			ret = append(ret, command)
		}
	}
	return ret
}
