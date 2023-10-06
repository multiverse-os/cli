package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	cli "github.com/multiverse-os/cli"

	color "github.com/multiverse-os/cli/terminal/ansi/color"
	rectangles "github.com/multiverse-os/cli/terminal/loading/bars/rectangles"
	circle "github.com/multiverse-os/cli/terminal/loading/spinners/circle"
)

func randomWait() {
	time.Sleep(time.Duration(rand.Intn(2)+2) * time.Second)
}

func main() {
	cmd, initErrors := cli.New(cli.App{
		Name:        "dev-cli",
		Description: "an example cli application for scripts and full-featured applications",
		Version:     cli.Version{Major: 0, Minor: 1, Patch: 1},
		GlobalFlags: cli.Flags(
			cli.Flag{
				Name:        "language",
				Alias:       "l",
				Default:     "en",
				Description: "Locale used when executing the program",
			},
			cli.Flag{
				Category:    "Server",
				Name:        "port",
				Alias:       "p",
				Default:     "3000",
				Description: "Port the server will listen on",
			},
			cli.Flag{
				Category:    "Server",
				Name:        "address",
				Alias:       "a",
				Description: "Host address the server will listen on",
			},
			cli.Flag{
				Category:    "Server",
				Name:        "daemon",
				Alias:       "d",
				Description: "Daemonize the program when launching",
			},
		),
		Commands: cli.Commands(
			cli.Command{
				Name:        "list",
				Alias:       "l",
				Description: "complete a task on the list",
				Action: func(c *cli.Context) error {
					spinner := c.CLI.Spinner(circle.Animation)

					spinner.Start()
					spinner.Message("Water, Dirt & Grass")
					randomWait()
					spinner.Message("Trees, Debris & Hideouts")
					randomWait()
					spinner.Message("Wildlife, Werewolves & Bandits")
					randomWait()
					spinner.Message("Sounds of wildlife & trees waving in the wind")
					randomWait()
					spinner.Message("Hiding treasure in the haunted woods...")
					randomWait()
					spinner.Complete("Completed")
					randomWait()

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
						Flags: cli.Flags(
							cli.Flag{
								Name:        "test",
								Alias:       "t",
								Default:     "what",
								Description: "A test filter",
							},
						),
						Action: func(c *cli.Context) error {
							loadingBar := c.CLI.LoadingBar(rectangles.Animation)

							loadingBar.Start()
							for i := 0; i < 100; i++ {
								time.Sleep(time.Duration(rand.Intn(2)+3) * time.Second)
								if loadingBar.Increment(1) {
									break
								}
							}
							loadingBar.Status(color.Green("Completed!")).End()

							fmt.Printf(
								"how many flags does context have (%v)\n",
								len(c.Flags),
							)

							c.CLI.Log("=====================================================")
							c.CLI.Log("====> c.Flag(\"l\"):", c.Flag("l").String())
							c.CLI.Log("add a thing to the list")
							for _, command := range c.Commands {
								c.CLI.Log("=====================================================")
								c.CLI.Log("[COMMAND:" + command.Name + "]")
								for _, flag := range command.Flags {
									c.CLI.Log("  '==>[FLAG][NAME:" + flag.Name + "][VALUE:" + flag.String() + "][DEFAULT:" + flag.Default + "]")
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
				Name:        "show",
				Alias:       "sh",
				Description: "show and item in the list",
				Action: func(c *cli.Context) error {
					c.CLI.Log("example action")
					return nil
				},
			},
		),
		Actions: cli.Actions{
			OnStart: func(c *cli.Context) error {
				//c.CLI.Log("OnStart action")
				return nil
			},
			//Fallback: func(c *cli.Context) error {
			//  c.CLI.Log("Fallback action")
			//  return nil
			//},
			OnExit: func(c *cli.Context) error {
				//c.CLI.Log("OnExit action")
				//  c.CLI.Log("=====================================================")
				//  // TODO: Switch to only using these and document this log system in the
				//  // API better
				//  c.CLI.Log("Command.Name:         ", c.Command.Name)
				//  c.CLI.Log("flag count [ ", string(c.Command.Flags.Count()), "] :")
				//  c.CLI.Log("=====================================================")

				//  for _, command := range c.Commands {
				//  	c.CLI.Log("=====================================================")
				//  	c.CLI.Log("command:", command.Name)
				//    //c.CLI.Log("command:action= [", command.Action, "]")
				//  	for _, flag := range command.Flags {
				//  		c.CLI.Log("command:flag= [", command.Name, "][", flag.Name, "][", flag.String(), "]")
				//  	}
				//  }

				//  for _, flag := range c.Flags {
				//  	c.CLI.Log("=====================================================")
				//  	c.CLI.Log("flag.Name :       ", flag.Name)
				//  	c.CLI.Log("flag.Value:       ", flag.String())
				//  }
				//  c.CLI.Log("=====================================================")

				return nil
			},
		},
	})

	if len(initErrors) == 0 {
		cmd.Parse(os.Args).Execute()
	}
}
