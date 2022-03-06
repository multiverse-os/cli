package cli

import (
	"fmt"
	"strconv"
	"strings"

	data "github.com/multiverse-os/cli/data"
	ansi "github.com/multiverse-os/cli/terminal/ansi"
	template "github.com/multiverse-os/cli/terminal/template"
)

type VersionComponent int

const (
	Major VersionComponent = iota
	Minor
	Patch
)

type Version struct {
	Major int
	Minor int
	Patch int
	Build BuildInformation
}

func MarshalVersion(version string) Version {
	components := strings.Split(version, ".")
	if len(components) == 3 {
		major, err := strconv.Atoi(components[0])
		minor, err := strconv.Atoi(components[1])
		patch, err := strconv.Atoi(components[2])
		if err != nil {
			return Version{Major: major, Minor: minor, Patch: patch}
		}
	}
	return Version{Major: 0, Minor: 0, Patch: 0}
}

type BuildInformation struct {
	Source     string
	Commit     string
	Signature  string
	CompiledAt string
}

func (self Build) AddDeveloper(name, email string) {
	self.Developers = append(self.Developers, Developer{
		Name:  name,
		Email: email,
	})
}

func (self VersionComponent) String() string {
	switch self {
	case Major:
		return "Major"
	case Minor:
		return "Minor"
	default: // Patch
		return "Patch"
	}
}

func (self Version) ColorString() string {
	var colorVersion []string
	for _, versionComponent := range strings.Split(self.String(), ".") {
		if versionComponent == "0" {
			colorVersion = append(colorVersion, ansi.Light(ansi.SkyBlue(versionComponent)))
		} else {
			colorVersion = append(colorVersion, ansi.Bold(ansi.Purple(versionComponent)))
		}
	}
	return ansi.Light(ansi.Blue("[")) + 
         ansi.Light(ansi.Blue("v")) + 
         strings.Join(colorVersion, ansi.White(".")) +
         ansi.Light(ansi.Blue("]"))
}

func (self CLI) RenderVersionTemplate() error {
	err := template.StdOut(defaultVersionTemplate(), map[string]string{
		"header":  ansi.Bold(ansi.SkyBlue(self.Context.Commands.First().Name)),
		"version": self.Version.ColorString(),
	})
	//"build": table.New(BuildInformation{
	//	Source:     "n/a",
	//	Commit:     "n/a",
	//	Signature:  "n/a",
	//	CompiledAt: "n/a",
	//}).String(),
	if data.NotNil(err) {
		return err
	}
	return nil
}

func defaultVersionTemplate() string {
	return "{{.header}}" + ansi.SkyBlue(ansi.Light(" version ")) + "{{.version}}\n"
}

func (self Version) undefined() bool {
	return self.Major == 0 &&
         self.Minor == 0 &&
         self.Patch == 0
}

///////////////////////////////////////////////////////////////////////////////
// TODO: Add sorting
func (self Version) IsSame(v Version) bool {
	return self.Major == v.Major && 
         self.Minor == v.Minor &&
         self.Patch == v.Patch
}

func (self Version) IsOlderThan(v Version) bool {
	return self.Major < v.Major ||
		     (self.Major == v.Major && (self.Minor < v.Minor ||
         (self.Minor == v.Minor && self.Patch < v.Patch)))
}

func (self Version) IsNewerThan(v Version) bool {
	return self.Major > v.Major ||
		     (self.Major == v.Major && (self.Minor > v.Minor || 
         (self.Minor == v.Minor && self.Patch > v.Patch)))
}

func (self Version) String() string {
	return fmt.Sprintf("%v.%v.%v", self.Major, self.Minor, self.Patch)
}
