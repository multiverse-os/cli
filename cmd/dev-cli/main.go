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
		GlobalFlags: cli.Flags(
			cli.Flag{
				Name:        "lang",
				Alias:       "l",
				Default:     "english",
				Description: "Locale used when executing the program",
			},
			cli.Flag{
				Name:        "daemon",
				Alias:       "d",
				Default:     "false",
				Description: "Daemonize the program when launching",
			},
		),
		Commands: cli.Commands(
			cli.Command{
				Name:        "list",
				Alias:       "c",
				Description: "complete a task on the list",
        //Action: func(c *cli.Context) error {
        //  fmt.Println("list!")
        //  return nil
        //},
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
							fmt.Println("=====================================================")
							fmt.Println("====> c.Flag(\"lang\"):", c.Flag("lang").String())
							fmt.Println("add a thing to the list")
							for _, command := range c.CommandChain.Commands {
								fmt.Println("=====================================================")
								fmt.Println("[COMMAND:" + command.Name + "]")
								for _, flag := range command.Flags {
									fmt.Println("       `'==>[FLAG][NAME:" + flag.Name + "][VALUE:" + flag.Value + "][DEFAULT:" + flag.Default + "]")
								}
							}
							for flagName, flagValue := range c.Flags {
								fmt.Println("=====================================================")
								c.CLI.Log(cli.INFO, "flag.Name :       ", flagName)
								c.CLI.Log(cli.INFO, "flag.Value:       ", flagValue)
							}
							fmt.Println("=====================================================")

							return nil
						},
					},
				),
			},
			cli.Command{
				Name:        "add",
				Alias:       "a",
				Description: "add a task to the list",
        Action: func(c *cli.Context) error {
          fmt.Println("add")
          return nil
        },
			},
		),
    Actions: cli.Actions{
      Fallback: func(c *cli.Context) error {
        return nil
      },
      Global: func(c *cli.Context) error {
			  fmt.Println("=====================================================")
			  fmt.Println("====> c.Flag(\"lang\"):", c.Flag("lang").String())

			  fmt.Println("=====================================================")
			  c.CLI.Log(cli.INFO, "Command.Name:         ", c.Command.Name)
			  c.CLI.Log(cli.INFO, "flag count [ ", len(c.Command.Flags), "] :")
			  fmt.Println("=====================================================")

			  for _, command := range c.CommandChain.Commands {
			  	fmt.Println("=====================================================")
			  	fmt.Println("command:", command.Name)
          //fmt.Println("command:action= [", command.Action, "]")
			  	for _, flag := range command.Flags {
			  		fmt.Println("command:flag= [", command.Name, "][", flag.Name, "][", flag.Value, "]")
			  	}
			  }

			  for flagName, flagValue := range c.Flags {
			  	fmt.Println("=====================================================")
			  	c.CLI.Log(cli.INFO, "flag.Name :       ", flagName)
			  	c.CLI.Log(cli.INFO, "flag.Value:       ", flagValue)
			  }
			  fmt.Println("=====================================================")

			  return nil
		  },
    },
	})

	// NOTE: Has the ability output context and error, this enables developers to
	// handle their own routing or actions based on parsed context.
	// context, _ := cmd.Parse(os.Args)
	cmd.Parse(os.Args)
}
