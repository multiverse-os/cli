package argument

type Chain struct {
	Commands []*Command
}

func (self *Chain) Route(path []string) (*Command, bool) {
	cmd := &Command{}
	for index, command := range self.Commands {
		if command.Name == path[index] {
			if index == len(path) {
				return command, true
			} else {
				cmd = command
			}
		} else {
			return cmd, (len(cmd.Name) == 0)
		}
	}
	return nil, (len(cmd.Name) == 0)
}

func (self *Chain) AddCommand(command *Command) {
	self.Commands = append(self.Commands, command)
}

func (self *Chain) Path() (path []string) {
	for _, command := range self.Commands {
		path = append(path, command.Name)
	}
	return path
}
