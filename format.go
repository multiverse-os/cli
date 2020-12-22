package cli

import (
	"strings"
)

const prefixSize = 2
const tabSize = 4

func Prefix() string { return strings.Repeat(" ", 2) }
func Tab() string    { return strings.Repeat(" ", tabSize) }
