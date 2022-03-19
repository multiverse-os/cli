package cli_test

import (
  //"fmt"
  //"os"
  "testing"

  cli "github.com/multiverse-os/cli"
)

type testArg []string

func testArgs() []testArg {
  return  []testArg{
    testArg{"h"},
    testArg{"help"},
    testArg{"-h"},
    testArg{"-help"},
    testArg{"v"},
    testArg{"version"},
    testArg{"-v"},
    testArg{"-version"},
  }
}

func InitApp() cli.App {
  return cli.App{
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
              c.CLI.Log("example action")
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
      Fallback: func(c *cli.Context) error {
        //c.CLI.Log("Fallback action")
        return nil
      },
      OnExit: func(c *cli.Context) error {
        c.CLI.Log("OnExit action")
        return nil
      },
    },
  }
}

func Test_New(t *testing.T) {
  // NOTE: Test Empty App
  cli := cli.New()

  if cli == nil {
    t.Errorf("cli failed to create with empty cli.App, returned nil cli.CLI object")
  }
}


// TODO: Will need to define different cli objects using New() 

// TODO: Some obvio problems with these tests lol, wtf if you think this is
// normal omgffg
func Test_Parse(t *testing.T) {
  for _, arg := range testArgs() {
    cliContext, _ := cli.New().Parse(arg)
    if cliContext == nil {
      t.Errorf("cli failed to create with empty cli.App, returned nil cli.CLI object")
    }
  }
}
