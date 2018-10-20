package main

import (
	"fmt"
	"os"

	cli "github.com/multiverse-os/cli-framework"
)

func main() {
	cmd := cli.New(&cli.CLI{
		Name:    "Example Program",
		Version: cli.Version{Major: 0, Minor: 1, Patch: 1},
		Usage:   "make an explosive entrance",
		Action: func(c *cli.Context) error {
			fmt.Println("Example output for an action (or command)!")
			return nil
		},
	})

	cmd.Run(os.Args)
}
