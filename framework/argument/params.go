package argument

import (
	data "./data"
)

type Params struct {
	Position int
	Type     data.Type
	Value    []string
}
