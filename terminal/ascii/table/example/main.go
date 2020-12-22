package main

import (
	"../../table"
)

type House struct {
	Name  string
	Sigil string
	Motto string
}

func main() {
	Example()
}

func Example() {
	hs := []House{
		{"Stark", "direwolf", "Winter is coming"},
		{"Targaryen", "dragon", "Fire and Blood"},
		{"Lannister", "lion", "Hear Me Roar"},
	}

	// Output to stdout
	table.Output(hs)

	// Or just return table string and then do something
	s := table.Table(hs)
	_ = s
}
