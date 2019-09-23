package cli

import (
	"strconv"

	template "github.com/multiverse-os/cli/template"
	color "github.com/multiverse-os/cli/text/ansi/color"
	style "github.com/multiverse-os/cli/text/ansi/style"
	table "github.com/multiverse-os/cli/text/table"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

type BuildInformation struct {
	CompiledAt string
	Source     string
	Signature  string
}

func (self *CLI) RenderVersion() error {
	err := template.OutputStdOut(defaultVersionTemplate(), map[string]string{
		"header":  self.header(false),
		"version": self.Version.String(),
	})
	if err != nil {
		return err
	}
	return nil
}

func defaultVersionTemplate() string {

	buildInfo := []BuildInformation{
		{"n/a", "n/a", "n/a"},
	}

	return `
{{.header}}` +
		`  ` + color.White(style.Bold(`Version:`)) + ` {{.version}} ` +
		`
` + table.Table(buildInfo) + `

`
}

func (self Version) Undefined() bool {
	return (self.Major == 0 && self.Minor == 0 && self.Patch == 0)
}

func (self Version) OlderThan(v Version) bool {
	return (self.Major < v.Major || (self.Major == v.Major && self.Minor < v.Minor) ||
		(self.Major == v.Major && self.Minor == v.Minor && self.Patch < v.Patch))
}

func (self Version) NewerThan(v Version) bool {
	return (self.Major > v.Major || (self.Major == v.Major && self.Minor > v.Minor) ||
		(self.Major == v.Major && self.Minor == v.Minor && self.Patch > v.Patch))
}

func (self Version) String() (formattedString string) {
	if self.Major == 0 {
		formattedString = style.Thin(strconv.Itoa(self.Major))
	} else {
		formattedString = style.Bold(color.SkyBlue(strconv.Itoa(self.Major)))
	}
	formattedString += color.White(".")
	if self.Minor == 0 {
		formattedString += style.Thin(strconv.Itoa(self.Minor))
	} else {
		formattedString += style.Bold(color.SkyBlue(strconv.Itoa(self.Minor)))
	}
	formattedString += color.White(".")
	if self.Patch == 0 {
		formattedString += style.Thin(strconv.Itoa(self.Patch))
	} else {
		formattedString += style.Bold(color.SkyBlue(strconv.Itoa(self.Patch)))
	}
	return formattedString
}
