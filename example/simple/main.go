package main

import (
	"fmt"
	"os"

	cli "github.com/multiverse-os/cli"
)

func main() {
	cmd := cli.New(&cli.CLI{
		Name:        "example",
		Description: "make an explosive entrance",
		Version:     cli.Version{Major: 0, Minor: 1, Patch: 1},
		Flags: cli.Flags(
			cli.Flag{
				Name:         "lang",
				Alias:        "l",
				DefaultValue: "english",
				Description:  "Locale used when executing the program",
			},
		),
		Commands: cli.Commands(
			cli.Command{
				Name:        "list",
				Alias:       "c",
				Description: "complete a task on the list",
				Flags: cli.Flags(
					cli.Flag{
						Name:         "filter",
						Alias:        "f",
						DefaultValue: "all",
						Description:  "filter all the things",
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
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			cli.Command{
				Name:        "add",
				Alias:       "a",
				Description: "add a task to the list",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		),
		DefaultAction: func(c *cli.Context) error {
			fmt.Println("Command Path:         ", c.CommandPath)
			fmt.Println("Command Path Length:  ", len(c.CommandPath))
			fmt.Println("Command.Name:         ", c.Command.Name)
			fmt.Println("flags:")
			for _, flag := range c.Flags {
				fmt.Println("flag.Name :", flag.Name)
				fmt.Println("flag.Value:", flag.Value)
			}
			//c.CLI.Logger.Info("This should log to both stdout and file")
			//log.Log(log.INFO, "Test info JSON log").StdOut()
			//log.Log(log.INFO, "Test info JSON log").Format(log.JSON).File("./test.log")
			//log.Log(log.INFO, "Test JSON with values").ANSI().WithValue("test", "value").JSON().StdOut()
			//log.Log(log.DEBUG, "Lying for good things isnt bad.").Format(log.ANSI).WithValue("test", 10).StdOut()
			//log.Log(log.DEBUG, "Lying for good things isnt bad.").WithValue("pi", 3.14).WithError(errors.New("Test Error")).StdOut()
			//log.Log(log.DEBUG, "Lying for good things isnt bad.").WithValue("testBoolean", true).WithErrors([]error{errors.New("Test Error"), errors.New("Another test error")}).StdOut()
			return nil
		},
	})

	cmd.Run(os.Args)
}
