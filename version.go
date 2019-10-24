package cli

import (
	"strconv"

	template "github.com/multiverse-os/cli/framework/template"
	table "github.com/multiverse-os/cli/framework/text/table"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

func (self *CLI) renderVersion() error {
	err := template.OutputStdOut(defaultVersionTemplate(), map[string]string{
		"header":  self.asciiHeader("calvins"),
		"version": self.Version.String(),
	})
	if NotNil(err) {
		return err
	}
	return nil
}

// TODO: Table should have a generic table object we can use to fill in data
// when we dont have a struct so we dont have to resort to using anonymous
// structs like this if we dont want to
func defaultVersionTemplate() string {
	return "\n{{.header}}  " + color.White(style.Bold("Version:")) +
		" {{.version}} \n" + table.Table(struct {
		CompiledAt string
		Signature  string
		Source     string
	}{
		"n/a",
		"n/a",
		"n/a",
	}) + "\n"
}

func (self Version) undefined() bool {
	return (self.Major == 0 && self.Minor == 0 && self.Patch == 0)
}

// Public Methods
///////////////////////////////////////////////////////////////////////////////
func (self Version) OlderThan(v Version) bool {
	return (self.Major < v.Major || (self.Major == v.Major && self.Minor < v.Minor) ||
		(self.Major == v.Major && self.Minor == v.Minor && self.Patch < v.Patch))
}

func (self Version) NewerThan(v Version) bool {
	return (self.Major > v.Major || (self.Major == v.Major && self.Minor > v.Minor) ||
		(self.Major == v.Major && self.Minor == v.Minor && self.Patch > v.Patch))
}

func (self Version) String() string {
	return fmt.Sprintf("%v.%v.%v", self.Major, self.Minor, self.Path)
}
