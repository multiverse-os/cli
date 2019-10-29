package cli

// Mocking ////////////////////////////////////////////////////////////////////
func CommandTreeCLI() *CLI {
	return cli.New(&cli.CLI{
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
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			cli.Command{
				Name:        "export",
				Alias:       "a",
				Description: "add a task to the list",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		),
	})
}

// Tests //////////////////////////////////////////////////////////////////////
func TestParse(t *testing.T) {
	arguments := []string{"list", "add", "test"}
	cmd := CommandTreeCLI()

	expectedCommandPath = []string{"example", "list"}

	context := cmd.Parse(arguments)
	if actual != expected {
		t.Error("Test failed")
	}
}
