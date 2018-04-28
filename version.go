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

func String() string {
	return fmt.Sprintf(""+strconv.Itoa(Major)+"."+strconv.Itoa(Minor)+"."+strconv.Itoa(Path))
}
