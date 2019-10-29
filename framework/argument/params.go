package argument

import (
	data "github.com/multiverse-os/cli/framework/data"
)

type Params struct {
	Position int
	Type     data.Type
	Value    []string
}
