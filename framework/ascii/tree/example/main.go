package main

import (
	"fmt"

	tree "github.com/multiverse-os/cli/text/tree"
)

func main() {
	fmt.Println("text tree")
	fmt.Println("=========")

	t := tree.New()

	one := t.AddBranch("one")
	one.AddNode("subnode1").AddNode("subnode2")
	one.AddBranch("two").
		AddNode("subnode1").AddNode("subnode2"). // add some nodes
		AddBranch("three").                      // add a new sub-branch
		AddNode("subnode1").AddNode("subnode2")  // add some nodes too
	one.AddNode("subnode3")
	t.AddNode("outernode")

	fmt.Println(t.String())
}
