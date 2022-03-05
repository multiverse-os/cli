package main

import (
  "os"

	cli "github.com/multiverse-os/cli"
)

func main() {

	cmd := cli.New(&cli.App{
		Name:        "example",
		Description: "an example cli application for scripts and full-featured applications",
		Version:     cli.Version{Major: 0, Minor: 1, Patch: 1},
		GlobalFlags: cli.Flags(
			cli.Flag{
				Name:        "language",
				Alias:       "lang",
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
				Alias:       "l",
				Description: "complete a task on the list",
        Action: func(c *cli.Context) error {
          c.CLI.Log("list!")
          return nil
        },
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
							c.CLI.Log("=====================================================")
							c.CLI.Log("====> c.Flag(\"lang\"):", c.Flag("lang").String())
							c.CLI.Log("add a thing to the list")
							for _, command := range c.Commands {
								c.CLI.Log("=====================================================")
								c.CLI.Log("[COMMAND:" + command.Name + "]")
								for _, flag := range command.Flags {
									c.CLI.Log("       `'==>[FLAG][NAME:" + flag.Name + "][VALUE:" + flag.String() + "][DEFAULT:" + flag.Default + "]")
								}
							}
							for _, flag := range c.Flags {
								c.CLI.Log("=====================================================")
								c.CLI.Log("flag.Name :       ", flag.Name)
								c.CLI.Log("flag.Value:       ", flag.String())
							}
							c.CLI.Log("=====================================================")

							return nil
						},
					},
				),
			},
			cli.Command{
				Name:        "example",
				Alias:       "ex",
				Description: "example command",
        Action: func(c *cli.Context) error {
          c.CLI.Log("example action")
          return nil
        },
        Hooks: cli.Hooks{
          BeforeAction: func(c *cli.Context) error {
            return nil
          },
          AfterAction: func(c *cli.Context) error {
            return nil
          },
        },
			},
		),
    GlobalHooks: cli.Hooks{
      BeforeAction: func(c *cli.Context) error {
        return nil
      },
      AfterAction: func(c *cli.Context) error {
        return nil
      },
    },
    Actions: cli.Actions{
      Fallback: func(c *cli.Context) error {
        c.CLI.Log("fallback action")
        return nil
      },
      Global: func(c *cli.Context) error {
        c.CLI.Log("global action")
			  c.CLI.Log("=====================================================")
			  c.CLI.Log("====> c.Flag(\"lang\"):", c.Flag("lang").String())

			  c.CLI.Log("=====================================================")
        // TODO: Switch to only using these and document this log system in the
        // API better
			  c.CLI.Log("Command.Name:         ", c.Command.Name)
			  c.CLI.Log("flag count [ ", string(c.Command.Flags.Count()), "] :")
			  c.CLI.Log("=====================================================")

			  for _, command := range c.Commands {
			  	c.CLI.Log("=====================================================")
			  	c.CLI.Log("command:", command.Name)
          //c.CLI.Log("command:action= [", command.Action, "]")
			  	for _, flag := range command.Flags {
			  		c.CLI.Log("command:flag= [", command.Name, "][", flag.Name, "][", flag.String(), "]")
			  	}
			  }

			  for _, flag := range c.Flags {
			  	c.CLI.Log("=====================================================")
			  	c.CLI.Log("flag.Name :       ", flag.Name)
			  	c.CLI.Log("flag.Value:       ", flag.String())
			  }
			  c.CLI.Log("=====================================================")

			  return nil
		  },
    },
	})

	// NOTE: Has the ability output context and error, this enables developers to
	// handle their own routing or actions based on parsed context.
	// context, _ := cmd.Parse(os.Args)
  cmd.Parse(os.Args)
}
