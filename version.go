package cli

import (
	"fmt"
	"strconv"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

func (self Version) String() string {
	return fmt.Sprintf("" + strconv.Itoa(self.Major) + "." + strconv.Itoa(self.Minor) + "." + strconv.Itoa(self.Patch))
}

// TODO: Add is version >, and add code to indicate API breaking landmark updates
