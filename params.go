package cli

import (
	data "github.com/multiverse-os/cli/data"
)

type Params struct {
	Position int
	Type     data.Type
	Value    []string
}
