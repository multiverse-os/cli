package cli_test

import (
  "fmt"
  "reflect"
  "runtime"
  "strings"
  "testing"

  cli "github.com/multiverse-os/cli"
)

type parseTest struct {
  Args           []string
  ExpectedAction string
}

func parseTests() []parseTest {
  return  []parseTest{
    parseTest{
      Args: []string{"test-cli", "h"},
      ExpectedAction: "HelpCommand",
    }, 
    parseTest{
      Args: []string{"test-cli", "help"},
      ExpectedAction: "HelpCommand",
    }, 
    parseTest{
      Args: []string{"test-cli", "-h"},
      ExpectedAction: "RenderDefaultHelpTemplate",
    }, 
    parseTest{
      Args: []string{"test-cli", "--help"},
      ExpectedAction: "RenderDefaultHelpTemplate",
    }, 
    parseTest{
      Args: []string{"test-cli", "v"},
      ExpectedAction: "RenderDefaultVersionTemplate",
    }, 
    parseTest{
      Args: []string{"test-cli", "version"},
      ExpectedAction: "RenderDefaultVersionTemplate",
    }, 
    parseTest{
      Args: []string{"test-cli", "-v"},
      ExpectedAction: "RenderDefaultVersionTemplate",
    }, 
    parseTest{
      Args: []string{"test-cli", "--version"},
      ExpectedAction: "RenderDefaultVersionTemplate",
    }, 
    parseTest{
      Args: []string{"test-cli", "list", "--help"},
      ExpectedAction: "RenderDefaultHelpTemplate",
    }, 
    parseTest{
      Args: []string{"test-cli", "list", "help"},
      ExpectedAction: "HelpCommand",
    }, 
  }
}

func initTestApp() cli.App {
  return cli.App{
    Name:        "test-cli",
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
  _, initErrors := cli.New()

  if len(initErrors) != 0 {
    t.Errorf("want (0) errors, got (%v)", len(initErrors))
  }
}


// TODO: Will need to define different cli objects using New() 
func Test_Parse(t *testing.T) {
  for _, parseTest := range parseTests() {
    // NOTE: Empty New() used, so it should only have 1 action 
    cmd, _ := cli.New(initTestApp())
    cmd.Parse(parseTest.Args)
    fmt.Printf("cli args (%v)", parseTest.Args)

    for _, action := range cmd.Context.Actions {
      actionName := strings.TrimPrefix(
        // TODO: see if you can check against The valueOf, the Pointer
        // or anything but the name if at all possible
        runtime.FuncForPC(reflect.ValueOf(action).Pointer()).Name(),
        "github.com/multiverse-os/cli.",
      )
      if strings.Contains(actionName, "cli_test") {
        continue
      }

      fmt.Println("action name:", actionName)

      if actionName != parseTest.ExpectedAction {
        t.Errorf("want (%v) function, got (%v)", actionName, parseTest.ExpectedAction)
      }
    }
  }
}
