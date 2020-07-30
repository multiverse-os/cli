package cli

import (
	"fmt"
	"strconv"
	"strings"

	ansi "github.com/multiverse-os/ansi"
	data "github.com/multiverse-os/cli/framework/data"
	template "github.com/multiverse-os/cli/framework/template"
	//table "github.com/multiverse-os/cli/framework/text/table"
)

// Semantic Versioning
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
	Build *BuildInformation
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

func (self *Build) AddAuthor(name, email string) {
	self.Authors = append(self.Authors, Author{
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
	return ansi.Light(ansi.Blue("[")) + ansi.Light(ansi.Blue("v")) + strings.Join(colorVersion, ansi.White(".")) + ansi.Light(ansi.Blue("]"))
}

func (self *CLI) RenderVersionTemplate() error {
	err := template.OutputStdOut(defaultVersionTemplate(), map[string]string{
		"header":  ansi.Bold(ansi.SkyBlue(self.Name)),
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
	return (self.Major == 0 && self.Minor == 0 && self.Patch == 0)
}

func DefaultVersion() Version {
	return Version{Major: 0, Minor: 1, Patch: 0}
}

// Public Methods
///////////////////////////////////////////////////////////////////////////////
type Compare func(a, b int) bool

func (self Version) DefaultVersion() Version { return Version{Major: 0, Minor: 1, Patch: 0} }

func (self Version) CompareComponent(component VersionComponent, compare Compare, v Version) bool {
	switch component {
	case Major:
		return compare(self.Major, v.Major)
	case Minor:
		return compare(self.Minor, v.Minor)
	case Patch:
		return compare(self.Minor, v.Minor)
	default:
		return false
	}
}

func (self Version) IsSame(v Version) bool {
	return self.Major == v.Major && self.Minor == v.Minor && self.Patch == v.Patch
}

func (self Version) IsOlderThan(v Version) bool {
	return self.Major < v.Major ||
		(self.Major == v.Major && (self.Minor < v.Minor || (self.Minor == v.Minor && self.Patch < v.Patch)))
}

func (self Version) IsNewerThan(v Version) bool {
	return self.Major > v.Major ||
		(self.Major == v.Major && (self.Minor > v.Minor || (self.Minor == v.Minor && self.Patch > v.Patch)))
}

// TODO: Color should be done by splitting by '.' and joining with a newly
// colored. This is where coloring just based on regex would be nice. Just color
// all semantic versions. Color all IPs, etc. Or a specialized printer for
// various different types.
func (self Version) String() string {
	return fmt.Sprintf("%v.%v.%v", self.Major, self.Minor, self.Patch)
}
