cli-framework
===
cli-framework is a simple command line interface framework


## Getting Started
This app will run and show help text, but is not very useful. Let's give an
action to execute and some help documentation:

<!-- {
  "output": "boom! I say!"
} -->
``` go
package main

import (
  "fmt"
  "os"

  cli "github.com/hackwave/cli-framework"
)

func main() {
  cmd := cli.New()
  cmd.Name = "boom"
	// NOTE: This makes more sense as 'cmd' over 'app', because the application version (the backing library or protocol) is separate from the CLI version
  cmd.Version = &cli.Version{
  	Major: 0,
	  Minor: 1,
	  Patch: 0, 
  }
  cmd.Usage = "make an explosive entrance"
  cmd.Action = func(c *cli.Context) error {
    fmt.Println("boom! I say!")
    return nil
  }

  cmd.Run(os.Args)
}
```

Running this already gives you a ton of functionality, plus support for things
like subcommands and flags, which are covered below.

## Examples

Start by creating a directory named `greet`, and within it, add a file,
`greet.go` with the following code in it:

<!-- {
  "output": "Hello friend!"
} -->
``` go
package main

import (
  "fmt"
  "os"

  cli "github.com/hackwave/cli-framework"
)

func main() {
  cmd := cli.New()
  cmd.Name = "greet"
  cmd.Usage = "fight the loneliness!"
  cmd.Action = func(c *cli.Context) error {
    fmt.Println("Hello friend!")
    return nil
  }

  cmd.Run(os.Args)
}
```

Install our command to the `$GOPATH/bin` directory:

```
$ go install
```

Finally run our new command:

```
$ greet
Hello friend!
```

cli also generates neat help text:

```
$ greet help
NAME:
    greet - fight the loneliness!

USAGE:
    greet [global options] command [command options] [arguments...]

VERSION:
    0.0.0

COMMANDS:
    help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS
    --version Shows version information
```

### Arguments

You can lookup arguments by calling the `Args` function on `cli.Context`, e.g.:

<!-- {
  "output": "Hello \""
} -->
``` go
package main

import (
  "fmt"
  "os"

  cli "github.com/hackwave/cli-framework"
)

func main() {
  cmd := cli.New()
  cmd.Action = func(c *cli.Context) error {
    fmt.Printf("Hello %q", c.Args().Get(0))
    return nil
  }
  cmd.Run(os.Args)
}
```

### Flags

Setting and querying flags is simple.

<!-- {
  "output": "Hello Nefertiti"
} -->
``` go
package main

import (
  "fmt"
  "os"

  cli "github.com/hackwave/cli-framework"
)

func main() {
  cmd := cli.New()

  cmd.Flags = []cli.Flag {
    cli.StringFlag{
      Name: "lang",
      Value: "english",
      Usage: "language for the greeting",
    },
  }

  cmd.Action = func(c *cli.Context) error {
    name := "Nefertiti"
    if c.NArg() > 0 {
      name = c.Args().Get(0)
    }
    if c.String("lang") == "spanish" {
      fmt.Println("Hola", name)
    } else {
      fmt.Println("Hello", name)
    }
    return nil
  }

  cmd.Run(os.Args)
}
```

You can also set a destination variable for a flag, to which the content will be
scanned.

<!-- {
  "output": "Hello someone"
} -->
``` go
package main

import (
  "os"
  "fmt"

  cli "github.com/hackwave/cli-framework"
)

func main() {
  var language string

  cmd := cli.New()

  cmd.Flags = []cli.Flag {
    cli.StringFlag{
      Name:        "lang",
      Value:       "english",
      Usage:       "language for the greeting",
      Destination: &language,
    },
  }

  cmd.Action = func(c *cli.Context) error {
    name := "someone"
    if c.NArg() > 0 {
      name = c.Args()[0]
    }
    if language == "spanish" {
      fmt.Println("Hola", name)
    } else {
      fmt.Println("Hello", name)
    }
    return nil
  }

  cmd.Run(os.Args)
}
```

A full list of flags can be found within the source code under the file `flag.go`.

#### Placeholder Values

Sometimes it's useful to specify a flag's value within the usage string itself.
Such placeholders are indicated with back quotes.

For example this:

<!-- {
  "args": ["&#45;&#45;help"],
  "output": "&#45;&#45;config FILE, &#45;c FILE"
} -->
```go
package main

import (
  "os"

  cli "github.com/hackwave/cli"
)

func main() {
  cmd := cli.New()

  cmd.Flags = []cli.Flag{
    cli.StringFlag{
      Name:  "config, c",
      Usage: "Load configuration from `FILE`",
    },
  }

  cmd.Run(os.Args)
}
```

Will result in help output like:

```
--config FILE, -c FILE   Load configuration from FILE
```

Note that only the first placeholder is used. Subsequent back-quoted words will
be left as-is.

#### Alternate Names

You can set alternate (or short) names for flags by providing a comma-delimited
list for the `Name`. e.g.

<!-- {
  "args": ["&#45;&#45;help"],
  "output": "&#45;&#45;lang value, &#45;l value.*language for the greeting.*default: \"english\""
} -->
``` go
package main

import (
  "os"

  cli "github.com/hackwave/cli-framework"
)

func main() {
  cmd := cli.New()

  cmd.Flags = []cli.Flag {
    cli.StringFlag{
      Name: "lang, l",
      Value: "english",
      Usage: "language for the greeting",
    },
  }

  cmd.Run(os.Args)
}
```

That flag can then be set with `--lang spanish` or `-l spanish`. Note that
giving two different forms of the same flag in the same command invocation is an
error.

#### Ordering

Flags for the application and commands are shown in the order they are defined.
However, it's possible to sort them from outside this library by using `FlagsByName`
or `CommandsByName` with `sort`.

For example this:

<!-- {
  "args": ["&#45;&#45;help"],
  "output": "add a task to the list\n.*complete a task on the list\n.*\n\n.*\n.*Load configuration from FILE\n.*Language for the greeting.*"
} -->
``` go
package main

import (
  "os"
  "sort"

  cli "github.com/hackwave/cli-framework"
)

func main() {
  cmd := cli.New()

  cmd.Flags = []cli.Flag {
    cli.StringFlag{
      Name: "lang, l",
      Value: "english",
      Usage: "Language for the greeting",
    },
    cli.StringFlag{
      Name: "config, c",
      Usage: "Load configuration from `FILE`",
    },
  }

  cmd.Commands = []cli.Command{
    {
      Name:    "complete",
      Aliases: []string{"c"},
      Usage:   "complete a task on the list",
      Action:  func(c *cli.Context) error {
        return nil
      },
    },
    {
      Name:    "add",
      Aliases: []string{"a"},
      Usage:   "add a task to the list",
      Action:  func(c *cli.Context) error {
        return nil
      },
    },
  }

  sort.Sort(cli.FlagsByName(cmd.Flags))
  sort.Sort(cli.CommandsByName(cmd.Commands))

  cmd.Run(os.Args)
}
```

Will result in help output like:

```
--config FILE, -c FILE  Load configuration from FILE
--lang value, -l value  Language for the greeting (default: "english")
```

#### Values from the Environment

You can also have the default value set from the environment via `EnvVar`.  e.g.

<!-- {
  "args": ["&#45;&#45;help"],
  "output": "language for the greeting.*APP_LANG"
} -->
``` go
package main

import (
  "os"

  cli "github.com/hackwave/cli-framework"
)

func main() {
  cmd := cli.New()

  cmd.Flags = []cli.Flag {
    cli.StringFlag{
      Name: "lang, l",
      Value: "english",
      Usage: "language for the greeting",
      EnvVar: "APP_LANG",
    },
  }

  cmd.Run(os.Args)
}
```

The `EnvVar` may also be given as a comma-delimited "cascade", where the first
environment variable that resolves is used as the default.

<!-- {
  "args": ["&#45;&#45;help"],
  "output": "language for the greeting.*LEGACY_COMPAT_LANG.*APP_LANG.*LANG"
} -->
``` go
package main

import (
  "os"

  "github.com/urfave/cli"
)

func main() {
  cmd := cli.New()

  cmd.Flags = []cli.Flag {
    cli.StringFlag{
      Name: "lang, l",
      Value: "english",
      Usage: "language for the greeting",
      EnvVar: "LEGACY_COMPAT_LANG,APP_LANG,LANG",
    },
  }

  cmd.Run(os.Args)
}
```

#### Values from files

You can also have the default value set from file via `FilePath`.  e.g.

<!-- {
  "args": ["&#45;&#45;help"],
  "output": "password for the mysql database"
} -->
``` go
package main

import (
  "os"

  "github.com/urfave/cli"
)

func main() {
  cmd := cli.New()

  cmd.Flags = []cli.Flag {
    cli.StringFlag{
      Name: "password, p",
      Usage: "password for the mysql database",
      FilePath: "/etc/mysql/password",
    },
  }

  cmd.Run(os.Args)
}
```

Note that default values set from file (e.g. `FilePath`) take precedence over
default values set from the enviornment (e.g. `EnvVar`).

#### Values from alternate input sources (YAML, TOML, and others)

There is a separate package altsrc that adds support for getting flag values
from other file input sources.

Currently supported input source formats:
* YAML
* TOML

In order to get values for a flag from an alternate input source the following
code would be added to wrap an existing cli.Flag like below:

``` go
  altsrc.NewIntFlag(cli.IntFlag{Name: "test"})
```

Initialization must also occur for these flags. Below is an example initializing
getting data from a yaml file below.

``` go
  command.Before = altsrc.InitInputSourceWithContext(command.Flags, NewYamlSourceFromFlagFunc("load"))
```

The code above will use the "load" string as a flag name to get the file name of
a yaml file from the cli.Context.  It will then use that file name to initialize
the yaml input source for any flags that are defined on that command.  As a note
the "load" flag used would also have to be defined on the command flags in order
for this code snipped to work.

Currently only the aboved specified formats are supported but developers can
add support for other input sources by implementing the
altsrc.InputSourceContext for their given sources.

Here is a more complete sample of a command using YAML support:

<!-- {
  "args": ["test-cmd", "&#45;&#45;help"],
  "output": "&#45&#45;test value.*default: 0"
} -->
``` go
package notmain

import (
  "fmt"
  "os"

  "github.com/urfave/cli"
  "github.com/urfave/cli/altsrc"
)

func main() {
  cmd := cli.New()

  flags := []cli.Flag{
    altsrc.NewIntFlag(cli.IntFlag{Name: "test"}),
    cli.StringFlag{Name: "load"},
  }

  cmd.Action = func(c *cli.Context) error {
    fmt.Println("yaml ist rad")
    return nil
  }

  cmd.Before = altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("load"))
  cmd.Flags = flags

  cmd.Run(os.Args)
}
```

#### Precedence

The precedence for flag value sources is as follows (highest to lowest):

0. Command line flag value from user
0. Environment variable (if specified)
0. Configuration file (if specified)
0. Default defined on the flag

### Subcommands

Subcommands can be defined for a more git-like command line app.

<!-- {
  "args": ["template", "add"],
  "output": "new task template: .+"
} -->
```go
package main

import (
  "fmt"
  "os"

  "github.com/urfave/cli"
)

func main() {
  cmd := cli.New()

  cmd.Commands = []cli.Command{
    {
      Name:    "add",
      Aliases: []string{"a"},
      Usage:   "add a task to the list",
      Action:  func(c *cli.Context) error {
        fmt.Println("added task: ", c.Args().First())
        return nil
      },
    },
    {
      Name:    "complete",
      Aliases: []string{"c"},
      Usage:   "complete a task on the list",
      Action:  func(c *cli.Context) error {
        fmt.Println("completed task: ", c.Args().First())
        return nil
      },
    },
    {
      Name:        "template",
      Aliases:     []string{"t"},
      Usage:       "options for task templates",
      Subcommands: []cli.Command{
        {
          Name:  "add",
          Usage: "add a new template",
          Action: func(c *cli.Context) error {
            fmt.Println("new task template: ", c.Args().First())
            return nil
          },
        },
        {
          Name:  "remove",
          Usage: "remove an existing template",
          Action: func(c *cli.Context) error {
            fmt.Println("removed task template: ", c.Args().First())
            return nil
          },
        },
      },
    },
  }

  cmd.Run(os.Args)
}
```

### Subcommands categories

For additional organization in apps that have many subcommands, you can
associate a category for each command to group them together in the help
output.

E.g.

```go
package main

import (
  "os"

  "github.com/urfave/cli"
)

func main() {
  cmd := cli.New()

  cmd.Commands = []cli.Command{
    {
      Name: "noop",
    },
    {
      Name:     "add",
      Category: "Template actions",
    },
    {
      Name:     "remove",
      Category: "Template actions",
    },
  }

  cmd.Run(os.Args)
}
```

Will include:

```
COMMANDS:
    noop

  Template actions:
    add
    remove
```

### Exit code

Calling `App.Run` will not automatically call `os.Exit`, which means that by
default the exit code will "fall through" to being `0`.  An explicit exit code
may be set by returning a non-nil error that fulfills `cli.ExitCoder`, *or* a
`cli.MultiError` that includes an error that fulfills `cli.ExitCoder`, e.g.:

``` go
package main

import (
  "os"

  "github.com/urfave/cli"
)

func main() {
  cmd := cli.New()
  cmd.Flags = []cli.Flag{
    cli.BoolTFlag{
      Name:  "ginger-crouton",
      Usage: "is it in the soup?",
    },
  }
  cmd.Action = func(ctx *cli.Context) error {
    if !ctx.Bool("ginger-crouton") {
      return cli.NewExitError("it is not in the soup", 86)
    }
    return nil
  }

  cmd.Run(os.Args)
}
```

### Bash Completion

You can enable completion commands by setting the `EnableBashCompletion`
flag on the `App` object.  By default, this setting will only auto-complete to
show an app's subcommands, but you can write your own completion methods for
the App or its subcommands.

<!-- {
  "args": ["complete", "&#45;&#45;generate&#45;bash&#45;completion"],
  "output": "laundry"
} -->
``` go
package main

import (
  "fmt"
  "os"

  "github.com/urfave/cli"
)

func main() {
  tasks := []string{"cook", "clean", "laundry", "eat", "sleep", "code"}

  cmd := cli.New()
  cmd.EnableBashCompletion = true
  cmd.Commands = []cli.Command{
    {
      Name:  "complete",
      Aliases: []string{"c"},
      Usage: "complete a task on the list",
      Action: func(c *cli.Context) error {
         fmt.Println("completed task: ", c.Args().First())
         return nil
      },
      BashComplete: func(c *cli.Context) {
        // This will complete if no args are passed
        if c.NArg() > 0 {
          return
        }
        for _, t := range tasks {
          fmt.Println(t)
        }
      },
    },
  }

  cmd.Run(os.Args)
}
```

Source the `autocomplete/bash_autocomplete` file in your `.bashrc` file while
setting the `PROG` variable to the name of your program:

`PROG=myprogram source /.../cli/autocomplete/bash_autocomplete`

#### Distribution

Copy `autocomplete/bash_autocomplete` into `/etc/bash_completion.d/` and rename
it to the name of the program you wish to add autocomplete support for (or
automatically install it there if you are distributing a package). Don't forget
to source the file to make it active in the current shell.

```
sudo cp src/bash_autocomplete /etc/bash_completion.d/<myprogram>
source /etc/bash_completion.d/<myprogram>
```

Alternatively, you can just document that users should source the generic
`autocomplete/bash_autocomplete` in their bash configuration with `$PROG` set
to the name of their program (as above).

#### Customization

The default bash completion flag (`--generate-bash-completion`) is defined as
`cli.BashCompletionFlag`, and may be redefined if desired, e.g.:

<!-- {
  "args": ["&#45;&#45;compgen"],
  "output": "wat\nhelp\nh"
} -->
``` go
package main

import (
  "os"

  "github.com/urfave/cli"
)

func main() {
  cli.BashCompletionFlag = cli.BoolFlag{
    Name:   "compgen",
    Hidden: true,
  }

  cmd := cli.New()
  cmd.EnableBashCompletion = true
  cmd.Commands = []cli.Command{
    {
      Name: "wat",
    },
  }
  cmd.Run(os.Args)
}
```

### Version Flag

The default version flag (`-v/--version`) is defined as `cli.VersionFlag`, which
is checked by the cli internals in order to print the `App.Version` via
`cli.VersionPrinter` and break execution.

#### Customization

The default flag may be customized to something other than `-v/--version` by
setting `cli.VersionFlag`, e.g.:

<!-- {
  "args": ["&#45;&#45print-version"],
  "output": "partay version 19\\.99\\.0"
} -->
``` go
package main

import (
  "os"

  "github.com/urfave/cli"
)

func main() {
  cli.VersionFlag = cli.BoolFlag{
    Name: "print-version, V",
    Usage: "print only the version",
  }

  cmd := cli.New()
  cmd.Name = "partay"
  cmd.Version = "19.99.0"
  cmd.Run(os.Args)
}
```

Alternatively, the version printer at `cli.VersionPrinter` may be overridden, e.g.:

<!-- {
  "args": ["&#45;&#45version"],
  "output": "version=19\\.99\\.0 revision=fafafaf"
} -->
``` go
package main

import (
  "fmt"
  "os"

  "github.com/urfave/cli"
)

var (
  Revision = "fafafaf"
)

func main() {
  cli.VersionPrinter = func(c *cli.Context) {
    fmt.Printf("version=%s revision=%s\n", c.App.Version, Revision)
  }

  cmd := cli.New()
  cmd.Name = "partay"
  cmd.Version = "19.99.0"
  cmd.Run(os.Args)
}
```

#### Full API Example

**Notice**: This is a contrived (functioning) example meant strictly for API
demonstration purposes.  Use of one's imagination is encouraged.

<!-- {
  "output": "made it!\nPhew!"
} -->
``` go
package main

import (
  "errors"
  "flag"
  "fmt"
  "io"
  "io/ioutil"
  "os"
  "time"

  "github.com/urfave/cli"
)

func init() {
  cli.AppHelpTemplate += "\nCUSTOMIZED: you bet ur muffins\n"
  cli.CommandHelpTemplate += "\nYMMV\n"
  cli.SubcommandHelpTemplate += "\nor something\n"

  cli.HelpFlag = cli.BoolFlag{Name: "halp"}
  cli.BashCompletionFlag = cli.BoolFlag{Name: "compgen", Hidden: true}
  cli.VersionFlag = cli.BoolFlag{Name: "print-version, V"}

  cli.HelpPrinter = func(w io.Writer, templ string, data interface{}) {
    fmt.Fprintf(w, "best of luck to you\n")
  }
  cli.VersionPrinter = func(c *cli.Context) {
    fmt.Fprintf(c.App.Writer, "version=%s\n", c.App.Version)
  }
  cli.OsExiter = func(c int) {
    fmt.Fprintf(cli.ErrWriter, "refusing to exit %d\n", c)
  }
  cli.ErrWriter = ioutil.Discard
  cli.FlagStringer = func(fl cli.Flag) string {
    return fmt.Sprintf("\t\t%s", fl.GetName())
  }
}

type hexWriter struct{}

func (w *hexWriter) Write(p []byte) (int, error) {
  for _, b := range p {
    fmt.Printf("%x", b)
  }
  fmt.Printf("\n")

  return len(p), nil
}

type genericType struct{
  s string
}

func (g *genericType) Set(value string) error {
  g.s = value
  return nil
}

func (g *genericType) String() string {
  return g.s
}

func main() {
  cmd := cli.New()
  cmd.Name = "kənˈtrīv"
  cmd.Version = "19.99.0"
  cmd.Compiled = time.Now()
  cmd.HelpName = "contrive"
  cmd.Usage = "demonstrate available API"
  cmd.UsageText = "contrive - demonstrating the available API"
  cmd.ArgsUsage = "[args and such]"
  cmd.Commands = []cli.Command{
    cli.Command{
      Name:        "doo",
      Aliases:     []string{"do"},
      Category:    "motion",
      Usage:       "do the doo",
      UsageText:   "doo - does the dooing",
      Description: "no really, there is a lot of dooing to be done",
      ArgsUsage:   "[arrgh]",
      Flags: []cli.Flag{
        cli.BoolFlag{Name: "forever, forevvarr"},
      },
      Subcommands: cli.Commands{
        cli.Command{
          Name:   "wop",
          Action: wopAction,
        },
      },
      SkipFlagParsing: false,
      HideHelp:        false,
      Hidden:          false,
      HelpName:        "doo!",
      BashComplete: func(c *cli.Context) {
        fmt.Fprintf(c.App.Writer, "--better\n")
      },
      Before: func(c *cli.Context) error {
        fmt.Fprintf(c.App.Writer, "brace for impact\n")
        return nil
      },
      After: func(c *cli.Context) error {
        fmt.Fprintf(c.App.Writer, "did we lose anyone?\n")
        return nil
      },
      Action: func(c *cli.Context) error {
        c.Command.FullName()
        c.Command.HasName("wop")
        c.Command.Names()
        c.Command.VisibleFlags()
        fmt.Fprintf(c.App.Writer, "dodododododoodododddooooododododooo\n")
        if c.Bool("forever") {
          c.Command.Run(c)
        }
        return nil
      },
      OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
        fmt.Fprintf(c.App.Writer, "for shame\n")
        return err
      },
    },
  }
  cmd.Flags = []cli.Flag{
    cli.BoolFlag{Name: "fancy"},
    cli.BoolTFlag{Name: "fancier"},
    cli.DurationFlag{Name: "howlong, H", Value: time.Second * 3},
    cli.Float64Flag{Name: "howmuch"},
    cli.GenericFlag{Name: "wat", Value: &genericType{}},
    cli.Int64Flag{Name: "longdistance"},
    cli.Int64SliceFlag{Name: "intervals"},
    cli.IntFlag{Name: "distance"},
    cli.IntSliceFlag{Name: "times"},
    cli.StringFlag{Name: "dance-move, d"},
    cli.StringSliceFlag{Name: "names, N"},
    cli.UintFlag{Name: "age"},
    cli.Uint64Flag{Name: "bigage"},
  }
  cmd.EnableBashCompletion = true
  cmd.HideHelp = false
  cmd.HideVersion = false
  cmd.BashComplete = func(c *cli.Context) {
    fmt.Fprintf(c.App.Writer, "lipstick\nkiss\nme\nlipstick\nringo\n")
  }
  cmd.Before = func(c *cli.Context) error {
    fmt.Fprintf(c.App.Writer, "HEEEERE GOES\n")
    return nil
  }
  cmd.After = func(c *cli.Context) error {
    fmt.Fprintf(c.App.Writer, "Phew!\n")
    return nil
  }
  cmd.CommandNotFound = func(c *cli.Context, command string) {
    fmt.Fprintf(c.App.Writer, "Thar be no %q here.\n", command)
  }
  cmd.OnUsageError = func(c *cli.Context, err error, isSubcommand bool) error {
    if isSubcommand {
      return err
    }

    fmt.Fprintf(c.App.Writer, "WRONG: %#v\n", err)
    return nil
  }
  cmd.Action = func(c *cli.Context) error {
    cli.DefaultAppComplete(c)
    cli.HandleExitCoder(errors.New("not an exit coder, though"))
    cli.ShowAppHelp(c)
    cli.ShowCommandCompletions(c, "nope")
    cli.ShowCommandHelp(c, "also-nope")
    cli.ShowCompletions(c)
    cli.ShowSubcommandHelp(c)
    cli.ShowVersion(c)

    categories := c.App.Categories()
    categories.AddCommand("sounds", cli.Command{
      Name: "bloop",
    })

    for _, category := range c.App.Categories() {
      fmt.Fprintf(c.App.Writer, "%s\n", category.Name)
      fmt.Fprintf(c.App.Writer, "%#v\n", category.Commands)
      fmt.Fprintf(c.App.Writer, "%#v\n", category.VisibleCommands())
    }

    fmt.Printf("%#v\n", c.App.Command("doo"))
    if c.Bool("infinite") {
      c.App.Run([]string{"cmd", "doo", "wop"})
    }

    if c.Bool("forevar") {
      c.App.RunAsSubcommand(c)
    }
    c.App.Setup()
    fmt.Printf("%#v\n", c.App.VisibleCategories())
    fmt.Printf("%#v\n", c.App.VisibleCommands())
    fmt.Printf("%#v\n", c.App.VisibleFlags())

    fmt.Printf("%#v\n", c.Args().First())
    if len(c.Args()) > 0 {
      fmt.Printf("%#v\n", c.Args()[1])
    }
    fmt.Printf("%#v\n", c.Args().Present())
    fmt.Printf("%#v\n", c.Args().Tail())

    set := flag.NewFlagSet("contrive", 0)
    nc := cli.NewContext(c.App, set, c)

    fmt.Printf("%#v\n", nc.Args())
    fmt.Printf("%#v\n", nc.Bool("nope"))
    fmt.Printf("%#v\n", nc.BoolT("nerp"))
    fmt.Printf("%#v\n", nc.Duration("howlong"))
    fmt.Printf("%#v\n", nc.Float64("hay"))
    fmt.Printf("%#v\n", nc.Generic("bloop"))
    fmt.Printf("%#v\n", nc.Int64("bonk"))
    fmt.Printf("%#v\n", nc.Int64Slice("burnks"))
    fmt.Printf("%#v\n", nc.Int("bips"))
    fmt.Printf("%#v\n", nc.IntSlice("blups"))
    fmt.Printf("%#v\n", nc.String("snurt"))
    fmt.Printf("%#v\n", nc.StringSlice("snurkles"))
    fmt.Printf("%#v\n", nc.Uint("flub"))
    fmt.Printf("%#v\n", nc.Uint64("florb"))
    fmt.Printf("%#v\n", nc.GlobalBool("global-nope"))
    fmt.Printf("%#v\n", nc.GlobalBoolT("global-nerp"))
    fmt.Printf("%#v\n", nc.GlobalDuration("global-howlong"))
    fmt.Printf("%#v\n", nc.GlobalFloat64("global-hay"))
    fmt.Printf("%#v\n", nc.GlobalGeneric("global-bloop"))
    fmt.Printf("%#v\n", nc.GlobalInt("global-bips"))
    fmt.Printf("%#v\n", nc.GlobalIntSlice("global-blups"))
    fmt.Printf("%#v\n", nc.GlobalString("global-snurt"))
    fmt.Printf("%#v\n", nc.GlobalStringSlice("global-snurkles"))

    fmt.Printf("%#v\n", nc.FlagNames())
    fmt.Printf("%#v\n", nc.GlobalFlagNames())
    fmt.Printf("%#v\n", nc.GlobalIsSet("wat"))
    fmt.Printf("%#v\n", nc.GlobalSet("wat", "nope"))
    fmt.Printf("%#v\n", nc.NArg())
    fmt.Printf("%#v\n", nc.NumFlags())
    fmt.Printf("%#v\n", nc.Parent())

    nc.Set("wat", "also-nope")

    ec := cli.NewExitError("ohwell", 86)
    fmt.Fprintf(c.App.Writer, "%d", ec.ExitCode())
    fmt.Printf("made it!\n")
    return ec
  }

  if os.Getenv("HEXY") != "" {
    cmd.Writer = &hexWriter{}
    cmd.ErrWriter = &hexWriter{}
  }

  cmd.Metadata = map[string]interface{}{
    "layers":     "many",
    "explicable": false,
    "whatever-values": 19.99,
  }

  cmd.Run(os.Args)
}

func wopAction(c *cli.Context) error {
  fmt.Fprintf(c.App.Writer, ":wave: over here, eh\n")
  return nil
}
```

### Combining short Bool options

Traditional use of boolean options using their shortnames look like this:
```
# cmd foobar -s -o
```

Suppose you want users to be able to combine your bool options with their shortname.  This
can be done using the **UseShortOptionHandling** bool in your commands.  Suppose your program
has a two bool flags such as *serve* and *option* with the short options of *-o* and
*-s* respectively. With **UseShortOptionHandling** set to *true*, a user can use a syntax
like:
```
# cmd foobar -so
```

If you enable the **UseShortOptionHandling*, then you must not use any flags that have a single
leading *-* or this will result in failures.  For example, **-option** can no longer be used.  Flags
with two leading dashes (such as **--options**) are still valid.