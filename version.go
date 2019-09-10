package cli

import (
	"fmt"
	"strconv"

	color "github.com/multiverse-os/cli/text/ansi/color"
	style "github.com/multiverse-os/cli/text/ansi/style"
)

type Version struct {
	Major int
	Minor int
	Patch int
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

func (self Version) String() string {
	return fmt.Sprintf("" + strconv.Itoa(self.Major) + "." + strconv.Itoa(self.Minor) + "." + strconv.Itoa(self.Patch))
}

func (self Version) StringWithANSI() (formattedString string) {
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
