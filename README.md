<img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4"  width="256px" height="256px" align="right" alt="Multiverse OS Logo">

## Multiverse: command-line interface 'cli' Go framework
**URL** [multiverse-os.org](https://multiverse-os.org)

The `cli-framework` aims to provide a consistent, security focused, internationalization (from the start), and other modern CLI features in such a way that it can be easily used in any software from script to full applications weather system service or web application. In addition to allowing interdepenent access to modular subpackages that can be used individually without being forced to include the entire `cli-framework`. 

**Multiverse OS Core Go Library**
The Multiverse `cli-framework` is designed to meet the requirements of Multiverse OS system programs, the result is a command-line interface framework where security and simplicity are development priorities. The framework includes all the necessary components for a high quality system application, featuring ANSI coloring, full VT-100 terminal control, variety of data input types, input validation, both `CLI` and `Sh` interfaces, daemonization, PID handling, custom CLI templates, basic OS interaction (for example, generic access to default cache and applcation directories provided by the OS), logging, localization, and ASCII graphs, QR codes, banners, and tables to render data in a human readable format. However unlike other libraries, each of these features beyond the most basic functionality needed for a CLI program are provided as sublibraries enabling devlepers to easily mix-and-match with their favorite libraries, or selectively use CLI framework components without creating bloated software. 

## Quick Start: the simplest example
The following command-line tool CLI application will run the `Action`. Unlessthe two default flags/commands: (1) **Help** accessible by the flag `--help` or `-h` or by the command `help` or `h`. (2) **Version** accessible by the flag `--version` or `-v` or by the command `version` or `v` which simply displays the version. 

``` go
package main

import (
  "fmt"
  "os"

  cli "github.com/multiverse-os/cli-framework"
)

func main() {
	// NOTE: This makes more sense as 'cmd' over 'app', because the application
  	// version (the backing library or protocol) is separate from the CLI version.
	cmd := cli.New(&cli.CLI{
		Name:    "Example",
		Version: cli.Version{Major: 0, Minor: 1, Patch: 1},
		Usage: "make an explosive entrance",
		Action: func(c *cli.Context) error {
			fmt.Println("Example output in response to a command (action)")
			return nil
			},
		})
		
		cmd.Run(os.Args)
}
```

Defining an action allows the developer to override the default action which
would be to display the help text. Action can be omitted to display help when
defining commands, subcommands, and flags which are defined below.

The current default output of the simplest configurations will generate this:

```
   _                 _  _ 
  |_| ___  ___  ___ | ||_|
  | || . ||___||  _|| || |
  |_||  _|     |___||_||_|
     |_|               [v0.1.0]
  Usage:
    ip-cli [command]
  
  Available Commands:
    help       Display help text, specify a command for in depth command help
    version    Display version, and compiler information
  
  Flags:
    -h, --help      help for ip-cli
        --version   version for ip-cli

```

## Examples
Below you will find a collection of examples to illustrate the various ways
the `cli-framework` can be used. 

### Arguments

Below is the simpliest initalization of the `cli` framework. It opts out of
commands, flags, and just passes down your arguments to the default action. It
still provides version, help output, and executable release information like
developer signature (this may move to a different library eventually). You can 
lookup arguments by calling the `Args` function on `cli.Context`, e.g.:

```go
package main

import (
  "fmt"
  "os"

  cli "github.com/multiverse-os/cli-framework"
)

func main() {
  cmd := cli.New(&cli.CLI{
    Action: func(c *cli.Context) error {
      fmt.Printf("Hello %q", c.Args().Get(0))
      return nil
    },
  })

  cmd.Run(os.Args)
}
```

### Flags

Setting and querying flags is simple.

```go
package main

import (
  "fmt"
  "os"

  cli "github.com/hackwave/cli-framework"
)

func main() {
  cmd := cli.New(&cli.CLI{
  Flags: []cli.Flag {
    cli.StringFlag{
      Name:    "lang",
      Aliases: []string{"l"},
      Value:   "english",
      Usage:   "language for the greeting",
    },
  },
  Action: func(c *cli.Context) error {
    name := "Nefertiti"
    if c.NArg() > 0 {
      name = c.Args().Get(0)
    }
    if c.Flags["lang"] == "spanish" {
      fmt.Println("Hola", name)
    } else {
      fmt.Println("Hello", name)
    }
    return nil
  })

  cmd.Run(os.Args)
}
```
