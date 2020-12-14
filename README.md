<img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4"  width="256px" height="256px" align="right" alt="Multiverse OS Logo">

▒█▀▀█ ▒█░░░ ▀█▀ <br/>
▒█░░░ ▒█░░░ ▒█░ <br/>
▒█▄▄█ ▒█▄▄█ ▄█▄ <br/>

**URL** [multiverse-os.org](https://multiverse-os.org)
The `cli-framework` aims to provide a consistent, security focused, framework for creating command-line tools from the standard command-processor (commands, flags, parameters), shell interfaces, and background daemons.

The framework is specifically designed to have an incredibly light code footprint, with each of the features divded into subpackages enabling developers to select just the components they need; so the framework can be suited for simple scripts to full applications. 

The framework design and how developers interact wtih it is inspired by web application frameworks. 

Defining commands should be familiar as possible, using design pattern's established in Golang web application frameworks. Providing similar features: middleware, templating, and importantly security related functionality like user input validation. By prioritizing security related functionality, we establish a minimum baseline for applications so less people roll their own, and more people collaborate on important features relevant to all applications. 



#### Multiverse OS Core Framework
`cli` framework is designed to meet the requirements of Multiverse OS system applications; since this powers the low-level interface of each core application, Multiverse developers understand the importance of opting for simplicity, while still trying to provide a complete and intuitive user experience. *This is not production ready, it just reached v0.1.0, it currently does not provide validation or have adequate sanitization for both input, but also output printed to Terminal.*

**Features** 

  * **Full VT100 support** providing ANSI coloring, cursor, and terminal control. (A grid system is being developed, improvements to color to make it even easier to use, and CSS styling are planned features). 
  
  * **Built-ins for user input** including secure password input, list/menu, multiselect, shell, and input validaiton for all basic types. (Not fully implemented)
  
  * **Support for both command processor (flags, commands, and subcommands), and shell style `cli` interfaces**. In addition to interactive CLI tools, the Multiverse OS `cli` framework provides functionality for daemonization, PID handling, singals, and other basic functionality of a service. 
  
  * **Custom help, version, and shell output via basic templates** and soon formatting will extend to logging, human readable output.
  
  * **ASCII-based visualizations** in the form of **Tables**, **Graphs/Histograms**, **QR Codes**, **Banners**, and development continues to improve each visualization with default icon characters available by default in Debian linux.

  * **Animated loading bars, and spinners** with a variety of different animation styles, and an esay API to extend and implement custom loading animations.

  * **Localization support** extending passed just alternate strings, but correct formating of numbers, currency, and dates; inspired by rails.

  * *(Planned)* Middleware support to replace a overly complex hook before/after action. The goal is to simplify system apaplication development by using development patterns familiar to web developers to make the entire process more familiar and easier.

This library just laws down the foundation for user interface development, next Multiverse OS developers will be offering a`gui` framework to work in sync with this library to provide full functionality and user experience for system application.

#### Ontology of a command-line interfaces 
To simplify working with the framework, we need to define terms that will be used when interacting with the library.

The basic `cli` framework example is using commands and subcommands to route to a defined action (controller-like) having the paramter passed in a object that can be validated against any datatype and helpers to convert common input datatypes like `string`, `int`, `bool`, but also more specific types that are commonly used as input in CLI applications: `filename`, `filenames`, `filepath`, `url`, `ipv4`, `ipv6`, and `port`. 

```
 app-cli open /path/to/file.md
 \_____/ \_/  \______________/
    |     |          |
   app command   parameters
```

Command support is not just command and subcommand, but instead implemented as a
command tree, enabling recursive nesting of commands. 

```
  app-cli command subcommand
  app-cli command --command-flag=test subcommand --subcommand-flag test paramters
 --subcommand-flag flag2
```

Flags are assigned to the last command defined, and if the command does not have
the flag defined, it recursively checks the command's parent commands to provide
intuitive design over outputing and error and requiring the user to move the
command which can be tedious for novice users.


#### Flags & Type
Many other CLI frameworks have a bloated codebase, and overly complex
declaration because they have a special flag type for each datatype. For example
a `StringFlag` or `IntFlag`. 

A simple `Flag` object is used, and the type of data it takes in is not
declared. Instead when the developer is locating the flag data, it is done using
the typical function chain found in a lot of Multiverse OS libraries, and closed
with the type. 

So if the developer wants a `string` flag, the data is validated and the type is
guarnateed when pulling it out of the context. 

```
  langaugeFlag := context.Flag("langauge").String()
```

This simple change allows our delcarations to be simplified, consistent, and
importantly allowed us to reduce the codebase possibly more than 20%, while
maintaining similar or possibly more functionality than the more popular
implementation.

#### Multiverse OS Core Framework
`cli-framework` is designed to meet the requirements of Multiverse OS system applications; since this powers the low-level interface of each core application, Multiverse developers understand the importance of opting for simplicity, while still trying to provide a complete and intuitive user experience. *This is not production ready, it just reached v0.1.0, it currently does not provide validation or have adequate sanitization for both input, but also output printed to Terminal.*


**Features** 

  * **Full VT100 support** providing ANSI coloring, cursor, and terminal control. (A grid system is being developed, improvements to color to make it even easier to use, and CSS styling are planned features). 
  
  * **Built-ins for user input** including secure password input, list/menu, multiselect, shell, and input validaiton for all basic types. (Not fully implemented)
  
  * **Support for both command processor (flags, commands, and subcommands), and shell style `cli` interfaces**. In addition to interactive CLI tools, the Multiverse OS `cli` framework provides functionality for daemonization, PID handling, singals, and other basic functionality of a service. 
  
  * **Custom help, version, and shell output via basic templates** and soon formatting will extend to logging, human readable output.


## Quick Start: the simplest example
The following command-line tool CLI application will run the `Action`. Unlessthe two default flags/commands: (1) **Help** accessible by the flag `--help` or `-h` or by the command `help` or `h`. (2) **Version** accessible by the flag `--version` or `-v` or by the command `version` or `v` which simply displays the version. 

```go
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
		
		cmd.Parse(os.Args)
    // context := cmd.Parse(os.Args)
    // Can also obtain the context and do things with it 
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

  cmd.Parse(os.Args)
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
    cli.Flag{
      Name:    "language",
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
    if c.Flag("language").String() == "spanish" {
      fmt.Println("Hola", name)
    } else {
      fmt.Println("Hello", name)
    }
    return nil
  })

  cmd.Parse(os.Args)
}
```

### Commands and subcommands
Setting and querying flags is simple.

```
package main

import (
	"fmt"
	"os"

	cli "github.com/multiverse-os/cli"
	ip "github.com/multiverse-os/ip"
)

func main() {
	cmd := cli.New(&cli.CLI{
		Name:    "ip-cli",
		Version: cli.Version{Major: 0, Minor: 1, Patch: 0},
		Usage:   "Specify a command, and one ip address or more",
		Commands: []cli.Command{
			cli.Command{
				Name:    "lookup",
				Aliases: []string{"l"},
				Usage:   "look up information for a given ip address",
				Flags: []cli.Flag{
					cli.Flag{
						Name:  "ip",
						Usage: "address to lookup",
						Value: "8.8.8.8",
					},
				},
			},
			cli.Command{
				Name:    "draw",
				Aliases: []string{"d"},
				Usage:   "render line on globe showing connection",
				Flags: []cli.Flag{
					cli.Flag{
						Name:  "ip",
						Usage: "address to lookup",
						Value: "8.8.8.8",
					},
				},
			},
		},
		DefaultAction: func(context *cli.Context) error {
			fmt.Println("Drawing connection to 8.8.8.8")
			ip.DrawConnection("8.8.8.8")
			return nil
		},
	})

	cmd.Parse(os.Args)
}
```

#### Ongoing Design and API Issues
This will be our location to document issues arising as we determine limitations
of the given API and help us steer the project towards a finalized API we can
freeze.

```
 ls --flag=test --flag test
```

The lack of an "=" is exactly what indicates it is a boolean flag. We need to enforce "=" OR we have to enforce flag location OR enforce declare flag datatype on declaration.

This previously established aspect of the CLI api but it doesn't allow us to reasonably support bool flags. 

One solution is creating the `cli` object like a web router and attach flags and
commands like a router, using instead of GET/PUSH, using the flag type, or
command parameter type.

```
cli := cli.New(...)
// Command flags shouold probably just be defined insidie the command
// declaration for simplicity.
cli.Command("command", "subcommand").Flag(Bool, "/command/subcommand", "thing", []string{"f", "fi})
// OR
cli.Filename("flagname", "f", "fi")
// Commands could be where Filename is defining the parameter datatype
cli.Filename("/command/subcommand/:filename", controllerName())
// This is not looking great, but it could have one advantage is we could easily
// also provide a REST API, but we could do that without going this far into a
// URL style which simply doesnt reflect the expected input
```

The current flag system also does not support merging short flags for very short
command exeuction, for example: 

```
  ls -lah
```

So this needs to be included in the parser, because its very command and will
and should be expected.


