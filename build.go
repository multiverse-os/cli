package cli

import "time"

type Build struct {
	CompiledAt time.Time
	Source     string
	Commit     string
	Signature  string
	Developers []Developer
}

type Developer struct {
	PublicKey PublicKey
	Name      string
	Email     string
}
