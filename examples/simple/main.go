package main

import (
	"fmt"
	"os"

	cli "github.com/multiverse-os/cli"
)

func main() {
	cmd := cli.New(&cli.CLI{
		Name:        "example",
		Description: "an example cli application for scripts and full-featured applications",
		Version:     cli.Version{Major: 0, Minor: 1, Patch: 1},
		Flags: cli.Flags(
			cli.Flag{
				Name:        "lang",
				Alias:       "l",
				Default:     "english",
				Description: "Locale used when executing the program",
			},
		),
		Commands: cli.Commands(
			cli.Command{
				Name:        "list",
				Alias:       "c",
				Description: "complete a task on the list",
				Flags: cli.Flags(
					cli.Flag{
						Name:        "filter",
						Alias:       "f",
						Default:     "all",
						Description: "filter all the things",
					},
				),
				Subcommands: cli.Commands(
					cli.Command{
						Name:        "add",
						Description: "lists all of something",
						Action: func(c *cli.Context) error {
							fmt.Println("add a thing to the list")
							return nil
						},
					},
				),
			},
			cli.Command{
				Name:        "add",
				Alias:       "a",
				Description: "add a task to the list",
			},
		),
		DefaultAction: func(c *cli.Context) error {
			//c.CLI.Log(cli.INFO, "Command Path:         ", c.CommandPath)
			//c.CLI.Log(cli.INFO, "Command Path Length:  ", len(c.CommandPath))
			c.CLI.Log(cli.INFO, "Command.Name:         ", c.Command.Name)
			c.CLI.Log(cli.INFO, "flag count [ ", len(c.Command.Flags), "] :")
			for _, flag := range c.Flags {
				c.CLI.Log(cli.INFO, "flag.Name :       ", flag.Name)
				c.CLI.Log(cli.INFO, "flag.Value:       ", flag.Value)
			}
			return nil
		},
	})

	// NOTE: Has the ability output context and error, this enables developers to
	// handle their own routing or actions based on parsed context.
	cmd.Parse(os.Args)
}
