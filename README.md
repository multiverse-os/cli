<img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4"  width="256px" height="256px" align="right" alt="Multiverse OS Logo">


## Multiverse OS: Go 'cli' Framework

  ▒█▀▀█ ▒█░░░ ▀█▀  ┏━━┓                ┏┓
  ▒█░░░ ▒█░░░ ▒█░  ┃━┳╋┳┳━┓┏━━┳━┳┳┳┳━┳┳┫┣┓
  ▒█▄▄█ ▒█▄▄█ ▄█▄  ┃┏┛┃┏┫╋┗┫┃┃┃┻┫┃┃┃╋┃┏┫━┫
                   ┗┛ ┗┛┗━━┻┻┻┻━┻━━┻━┻┛┗┻┛ 
  
**URL** [multiverse-os.org](https://multiverse-os.org)

The `cli-framework` aims to provide a consistent, security focused, framework for creating command-line tools from the standard command-processor (commands, flags, arguments), shell interfaces, and background daemons. The framework design is directly inspired by feature complete web application frameworks like Ruby's `rails`, which translates to internationalization (in development), and other modern CLI features *(see below for full list)*. The framework is specifically designed to have an incredibly light code footprint, with each of the features divded into subpackages enabling developers to select just the components they need; so the framework can be suited for simple scripts to full applications. 


#### Ontology of a command-line interfaces 
To simplify working with the framework, we need to define terms that will be
used when interacting with the library. The software has a smaller code
footprint than other CLI frameworks, but it also is capable of offering more
complexity by combining logic to avoid unnecessary code repititon.

```
// For example a simple single command CLI tool allows the user to define a
command, and by default any *non-flag* values after the last (endpoint) command
is accepted as paramters, which is parsed into a comma separated slice. 
 app-cli open /path/to/file.md
 \_____/ \__/ \______________/
    |     |          |
   app  command  parameters
```

Other solutions have multiple struct types dedicated to each Flag data type
(bool, string, int,...), and instead of doing this we provided a generic
validation tool for flags given the desired type. This data can then be
validated, and if valid be output as the desired type. Doing it this way means
if we want to use the flag as an `int` value for one assignment, but since the
flag data itself is not typed, its typed to whatever we want, immediately after 
use it to assign a `string`. 

```
// The flags before commands are global flags within the application scope, and
// are listed when using the help command or flag (by default both are provided). 
// Flags following a command (even after the parameters) are first checked
// against the command flags listed when running help command/flag after a
// command: 

 app-cli command help
 \_____/  \___/  \__/
    |       |     |
   app    command subcommand 
```

The subcommand in this example is the endpoint (terminal, or edge) command. Our
coommands are recursively defined, generating a command tree. If the last
command when executing the the `app-cli` is not an endpoint command, meaning it
functions as a category for subcommands, it will by default output the help text
for that command. If it is the endpoint but it expected paramters, it will
output the help text for the endpoint command with an error by default. If no
parameters are required, and it is the endpoint command (edge of command tree)
it will execute as expected. 

If a developer overlooks the step of validating, and tries to output a flag as a
value not contained within the string, it will return the string value by
default (TODO: Decide if it should return an accompanying bool, or error
response with nil if it fails). 



```
app-cli command --command-flag=test subcommand --subcommand-flag test paramters
 --subcommand-flag flag2
```

Flags trailing will first check if the subcommand has the defined flag, then the
command, and finally the global scope. In this case, we just add it to the,
context as a global flag without an error by default.



#### Multiverse OS Core Framework
`cli-framework` is designed to meet the requirements of Multiverse OS system applications; since this powers the low-level interface of each core application, Multiverse developers understand the importance of opting for simplicity, while still trying to provide a complete and intuitive user experience. *This is not production ready, it just reached v0.1.0, it currently does not provide validation or have adequate sanitization for both input, but also output printed to Terminal.*


**Features** 

  * **Full VT100 support** providing ANSI coloring, cursor, and terminal control. (A grid system is being developed, improvements to color to make it even easier to use, and CSS styling are planned features). 
  
  * **Built-ins for user input** including secure password input, list/menu, multiselect, shell, and input validaiton for all basic types. (Not fully implemented)
  
  * **Support for both command processor (flags, commands, and subcommands), and shell style `cli` interfaces**. In addition to interactive CLI tools, the Multiverse OS `cli` framework provides functionality for daemonization, PID handling, singals, and other basic functionality of a service. 
  
  * **Custom help, version, and shell output via basic templates** and soon formatting will extend to logging, human readable output.
 app-cli --open-file /path/to/file.md
 \_____/  \_______/  \______________/
    |         |          |
   app      flag(global)  parameters
```


If a developer overlooks the step of validating,
and tries to output


#### Multiverse OS Core Framework
`cli-framework` is designed to meet the requirements of Multiverse OS system applications; since this powers the low-level interface of each core application, Multiverse developers understand the importance of opting for simplicity, while still trying to provide a complete and intuitive user experience. *This is not production ready, it just reached v0.1.0, it currently does not provide validation or have adequate sanitization for both input, but also output printed to Terminal.*


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
	// TODO: Ideally will just merge in or basically forward all existing `ip
	// commands` and use the command name `ip`, so basically it will function as a
	// way of adding in more consistent functionality (like adding JSON output to
	// every command, and providing more functionality, while keeping all the
	// original functionality and expected usage)

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

	cmd.Run(os.Args)
}

```
*(Add shell examples, loaders, various ASCII components, VT100, colors (brief), and link to various subpackage examples for full details)*
