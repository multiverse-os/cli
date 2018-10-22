package cli

type CommandCategories []CommandCategory

type CommandCategory struct {
	Name     string
	Hidden   bool
	Commands Commands
}

func (self CommandCategories) AddCommand(category string, command Command) CommandCategories {
	for _, category := range self {
		if commandCategory.Name == category {
			commandCategory.Commands = append(commandCategory.Commands, command)
			return self
		}
	}
	return append(self, &CommandCategory{Name: category, Commands: []Command{command}})
}

func (self CommandCategory) VisibleCommands() []Command {
	ret := []Command{}
	for _, command := range self.Commands {
		if !command.Hidden {
			ret = append(ret, command)
		}
	}
	return ret
}
