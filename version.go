package cli

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	data "github.com/multiverse-os/cli/framework/data"
	template "github.com/multiverse-os/cli/framework/template"
	color "github.com/multiverse-os/cli/framework/terminal/ansi/color"
	style "github.com/multiverse-os/cli/framework/terminal/ansi/style"
	banner "github.com/multiverse-os/cli/framework/text/banner"
	table "github.com/multiverse-os/cli/framework/text/table"
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
	Build *Build
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

//We can sign the build but unfortuatenly we can't sign the checksum because it would alter the checksum
// TODO: It would be great to impelement a middleware like system to
// make CLI programming similar to web programming. Reusing these conceepts
// should make it more familiar and easier to transpose code
// TODO: Provide a way to register an RSS feed that can be used for checking for
// updates.
type Build struct {
	Authors    []Author
	Source     string
	Commit     string
	Signature  string
	CompiledAt time.Time
}

type Author struct {
	Name  string
	Email string
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

func (self *CLI) RenderVersionTemplate() error {
	err := template.OutputStdOut(defaultVersionTemplate(), map[string]string{
		"header":  banner.New(self.Name).Font("big").String(),
		"version": self.Version.String(),
	})
	if data.NotNil(err) {
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
