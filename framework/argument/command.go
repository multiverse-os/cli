package argument

//
// Input Commands
///////////////////////////////////////////////////////////////////////////////
type Command struct {
	Definition interface{}
	Parent     *Command
	Name       string
	Flags      map[string]*Flag
}

func (self Command) Path() []string {
	route := []string{self.Name}
	for parent := self.Parent; parent != nil; parent = parent.Parent {
		route = append(route, parent.Name)
	}
	return route
}
