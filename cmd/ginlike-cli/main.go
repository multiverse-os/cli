package main

import (
  "fmt"

  cli "github.com/multiverse-os/cli"
)

func main() {
  fmt.Println("hi")

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
	})

	// NOTE: Has the ability output context and error, this enables developers to
	// handle their own routing or actions based on parsed context.
	// context, _ := cmd.Parse(os.Args)
	cmd.Parse(os.Args)
}
