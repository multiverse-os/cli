package cli

//"sync/atomic"

type CommandCategories map[string]CommandCategory

type CommandCategory struct {
	Name        string
	Description string
	Commands    Commands
	Hidden      bool
}

func InitCommandCategories(categoryName, categoryDescription string, commands Commands) CommandCategories {
	return CommandCategories{
		"key": CommandCategory{
			Name:        categoryName,
			Description: categoryDescription,
			Commands:    commands,
			Hidden:      false,
		},
	}
}

func (self CommandCategory) HasVisibleCommands() bool {
	return (len(self.VisibleCommands) > 0)
}

func (self CommandCategory) VisibleCommands() (commands []Command) {
	for _, command := range self.Commands {
		if !command.Hidden {
			commands = append(commands, command)
		}
	}
	return commands
}
