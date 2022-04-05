<img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4"  width="256px" height="256px" align="right" alt="Multiverse OS Logo">

▒█▀▀█ ▒█░░░ ▀█▀ <br/>
▒█░░░ ▒█░░░ ▒█░ <br/>
▒█▄▄█ ▒█▄▄█ ▄█▄ <br/>

**URL** [multiverse-os.org](https://multiverse-os.org)

*This library is still pre-alpha, but as of March 2022, it is very quickly
nearing alpha release. Which means we will freeze the API, and no changes will
be made to the API until at least first release candidate. So at that point you
can safely create applications and have no worry they will break from our
changes. We will soon be going over public and private functions, types and
fields to limit it exactly to what should be available for developers to use the
library and that will finalize the process of freezing the API.*

The `cli` framework aims to provide a security focused, and easy-to-use
toolbox for creating command-line interfaces for simple scripts, to full
featured TUI applications. Not just the standard command-processor model
(commands, flags, params) but also shell interfaces.

Example output from application using the `cli` framework: 

```
[DEBUG][Benchmark] benmarking argument parsing [ 4.1µs ]

　 █▀▀ ▀▀█▀▀ █▀▀█ █▀▀█ █▀▀ █  █  ▀  █▀▀█ █  █ █▀▀█ █▀▀█ █▀▀▄ 
　 ▀▀█   █   █▄▄█ █▄▄▀ ▀▀█ █▀▀█ ▀█▀ █  █ █▄▄█ █▄▄█ █▄▄▀ █  █ 
　 ▀▀▀   ▀   ▀  ▀ ▀ ▀▀ ▀▀▀ ▀  ▀ ▀▀▀ █▀▀▀ ▄▄▄█ ▀  ▀ ▀ ▀▀ ▀▀▀  0.1.1
  A command-line tool for controling the starshipyard server, scaffolding boilerplate code, and executing developer defined commands

  Usage starshipyard [options] [subcommand] [parameters]

  Subcommands
    console, c        Start the starship yard console interface
    new, n            Create a new starship project
    generate, g       Generate new go source code for models, controllers, and views
    server, s         Options for controlling starshipyard HTTP server
    version, v        outputs version

  Flags
   Global options
    -h, --help        outputs command and flag details

   Server options
    -e, --env         Specify the server environment [≅ development]
    -a, --address     Specify the address for the HTTP server to listen [≅ 0.0.0.0]
    -p, --port        Specify the listening port for the HTTP server [≅ 3000]
```

In contrast to other cli frameworks that try to be minimal as possible, so much so as to
leave out basic functionality like *stacked flags*, or *flag categories* (see git to see 
why developers might expect flag categories). `cli` is acheives a minimalist code-footprint
by encapsulating its features into optional sub-packages instead of leaving out
functionality that allows developers using the library to focus on what makes
their program unique instead of code common to other cli programs. The aim is
for `cli` to be reasonably small for use in simple scripts, and all the tools
needed for a modern full-featured cli application. 

This means providing not just secure user input, but a variety of rich user
inputs. And tools to actually create user interfaces beyond the help output.

`cli` provides multiple ways to generate ascii/text generation for tables when
and application needs to show the user data. And symbols to enrich the output
or create lists of data to present to the user. Banners, using figlet fonts, to
emphasize text, and what we use to improve the default help output. Sparkline 
graphs and other ascii/text based graphs for outputing small and large datasets
in a way that is meaningful to the user. And ANSI coloring and styling, 
provided individually in lean subpackages, or ANSI support for full VT100 TUI 
interfaces. 

Additionally user interfaces commonly need loading bars and spinners, so 
they are also included.

All of this, including the default help and version output is very easy to
customize or override completely. 

### CLI Framework
**Features** 
As this software is in its pre-alpha stages, not all the features below are
completed, some are complete, some are in-progress, and some are in planning
stages. 

* **Full VT100 support** providing ANSI coloring and styling through several 
sub-packages providing different levels of sophistication to provide
functionality for simple scripts with little overhead, or robust full CLI
applications with full SGR/CSI functionality with helpers, grid system, and
other features required for complete TUI applications.  

* **Sophisticated user input** including secure password input, list/menu, 
  multiselect, shell, and input validaiton for all basic types.

  * **Command-line interfaces with commands, subcommands, flags, and params**
  with full support for stacked flags, flag param separation using both " "
  and "=" for maximum compatibility.  

  * **Loading Bars & Spinners** easy to customize, and included as indepedent
  subpackages for minimal overhead. 

  * **ASCII/Text helpers** in the form of *Tables*, *Graphs/Histograms*, 
  *QR Codes*, *Banners* (using figlet fonts), symbol sets (using unicode) for 
  a variety of purposes. 

  * **Localization support**

### Example
  `dev-cli` is the example included with the package, and used in some of the test
  code.

  ```go
  package main

  import (
      "fmt"
      "os"
      "time"
      "math/rand"

      cli "github.com/multiverse-os/cli"

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
spinner := c.CLI.Spinner().Animation(circle.Animation)

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
loadingBar := c.CLI.LoadingBar().Animation(rectangles.Animation)



loadingBar.Start()
for i := 0; i < 100; i++ {
time.Sleep(time.Duration(rand.Intn(135)+22) * time.Millisecond)
if loadingBar.Increment(1) {
  break
}
}
loadingBar.End()

  //// NOTE: run code between two start and stop
  //c.Cli.LoadingBar(squards.Style).Stop()


  fmt.Printf("how many flags does context have (%v)\n", len(c.Flags))
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
```
