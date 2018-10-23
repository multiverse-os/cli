package cli

type CommandCategories map[string]CommandCategory

type CommandCategory struct {
	Name     string
	Commands Commands
	Hidden   bool
}

func (self CommandCategories) AddCommand(category string, command Command) CommandCategories {
	if category, ok := CommandCategories[category]; ok {
		return append(self, &CommandCategory{Name: category, Commands: []Command{command}})
	}
}

func (self CommandCategory) HasVisibleCommands() bool {
	return (len(self.Commands) > 0)
}

func (self CommandCategory) VisibleCommands() (commands []Command) {
	for _, command := range commands {
		if !command.Hidden {
			commands = append(commands, command)
		}
	}
	return commands
}
