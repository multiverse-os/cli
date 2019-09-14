package main

import (
	"errors"
	"fmt"
	"os"

	cli "github.com/multiverse-os/cli"
	log "github.com/multiverse-os/cli/log"
)

func main() {
	cmd := cli.New(&cli.CLI{
		Name:    "example",
		Version: cli.Version{Major: 0, Minor: 1, Patch: 1},
		Usage:   "make an explosive entrance",
		Flags: []cli.Flag{
			cli.Flag{
				Name:  "lang",
				Value: "english",
				Usage: "language for the greeting",
			},
		},
		Commands: []cli.Command{
			cli.Command{
				Name:    "complete",
				Aliases: []string{"c"},
				Usage:   "complete a task on the list",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			cli.Command{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		},
		DefaultAction: func(c *cli.Context) error {
			fmt.Println("command.Name:", c.Command.Name)
			fmt.Println("subcommand.Name:", c.Subcommand.Name)
			c.CLI.Logger.Info("This should log to both stdout and file")
			//log.Log(log.INFO, "Test info JSON log").StdOut()
			//log.Log(log.INFO, "Test info JSON log").Format(log.JSON).File("./test.log")
			//log.Log(log.INFO, "Test JSON with values").ANSI().WithValue("test", "value").JSON().StdOut()
			log.Log(log.DEBUG, "Lying for good things isnt bad.").Format(log.ANSI).WithValue("test", 10).StdOut()
			log.Log(log.DEBUG, "Lying for good things isnt bad.").WithValue("pi", 3.14).WithError(errors.New("Test Error")).StdOut()
			log.Log(log.DEBUG, "Lying for good things isnt bad.").WithValue("testBoolean", true).WithErrors([]error{errors.New("Test Error"), errors.New("Another test error")}).StdOut()
			return nil
		},
	})

	cmd.Run(os.Args)
}
