package cli

import (
	"fmt"
	"strconv"
	"strings"

	ansi "github.com/multiverse-os/ansi"
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

//func (b Build) AddDeveloper(name, email string) {
//	b.Developers = append(b.Developers, Developer{
//		Name:  name,
//		Email: email,
//	})
//}

func (vc VersionComponent) String() string {
	switch vc {
	case Major:
		return "Major"
	case Minor:
		return "Minor"
	default: // Patch
		return "Patch"
	}
}

func (v Version) ColorString() string {
	var colorVersion []string
	for _, versionComponent := range strings.Split(v.String(), ".") {
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

func (v Version) undefined() bool {
	return v.Major == 0 &&
		v.Minor == 0 &&
		v.Patch == 0
}

// /////////////////////////////////////////////////////////////////////////////
// TODO: Add sorting
func (v Version) IsSame(version Version) bool {
	return v.Major == version.Major &&
		v.Minor == version.Minor &&
		v.Patch == version.Patch
}

func (v Version) IsOlderThan(version Version) bool {
	return v.Major < version.Major ||
		(v.Major == version.Major && (v.Minor < version.Minor ||
			(v.Minor == version.Minor && v.Patch < version.Patch)))
}

func (v Version) IsNewerThan(version Version) bool {
	return v.Major > version.Major ||
		(v.Major == version.Major && (v.Minor > version.Minor ||
			(v.Minor == version.Minor && v.Patch > version.Patch)))
}

func (v Version) String() string {
	return fmt.Sprintf("%v.%v.%v", v.Major, v.Minor, v.Patch)
}
