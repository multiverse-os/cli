# Simple Logger
A very simple, optional and easy to override logger to work with the CLI
framework. It provides a simple interface to just write to screen with
function call as simple as `log.Info("info text")` or a more complex
`logger` that supports writing to a specified file. And the interface
is very simple and easy to override if a more complex logging system
is desired. 


## Getting Started
This minimalistic logging system has two approaches, (1) that requires the
initialization of a `Logger` object, that stores information about format,
outputs, and stores defined settings, (2) simple interface to log to standard
out of to a file with no previous initialization of variables or settings.

We will start by demostrating the later: 

```
package main 
import (
  log "github.com/multiverse-os/log"
)

func main() {
  log.Level(log.INFO).Message("Information to be sent to 'stdout'.").StdOut()
}
```

Our package provides a variety of aliases to provide intuitive as possible API:

```
log.Info("Information to be send to 'stdout'.").JSON().Output()

```

We are not limited to logging to standard out with this chaining setting
definitions:

```
log.Info("Information to be sent to
file.").XML().ToFile("~/.local/share/app/app.log")
```

But the simplest way is even simpler than chaining settings, we can send logs to
`stdout` with a single function call:

```
Info("Information to be sent to 'stdout'.")
```


## Using the 'Logger' Object

The `Logger` object provides a way to define settings once and store them inside
an object to be passed around or be an attribute in your application state so it
can be used throughout your software:

```
type State struct {
  Name string
  Logger log.Logger
}

func InitState() State {
  return State{
    Name: "app",
    // NewLogger expects: (appName, timestampResolution, logVerbosity)
    Logger: log.NewLogger("app", log.MINUTES, log.VERBOSE),
  }
}

func main() {
  state := InitState()
  state.Level(log.INFO).Message("Information to be sent to outputs")
}
```

**Alternative Initialization Methods For 'Logger' Object**
The following illustrates the variety of methods to initialize the `Logger`
object:

```
// NewSimpleLogger is init with: (App Name, Output Format, OutputToStdOut)
// The simple logger by default logs to the application data of the 
// user executing the binary: `~/.local/share/appName/appName.log`
logger := log.NewSimpleLogger("appName", log.JSON, false) 


// NewDefaultLogger is init with: (App Name, LogToDefaultUserFile, LogToStdOut)
// The default logger logs to default user location defined above as JSON 
// if the first bool is set to true, and logs to StdOut in DefaultWithANSI
// format if second bool is set to true.
logger := log.NewDefaultLogger("appName", true, true)
```

Then there is the function to create a `Logger` object with the maximum amount
of customization:

```
// The following is used to create a logger object with an arbitrary log file
logger := log.NewFileLogger("appName", log.SECONDS, log.VERY_VERBOSE, log.JSON, "/var/log/app-name.log")
```

After a `Logger` object is created, any output can be added after the fact,
`NewLogger` functions that add outputs by default are conviences to simplify
usage.

## Development
The biggest development priority currently is expanding the documentation, this
project is currently used as a sub-component of the Multiverse OS cli-framework,
used in all of Multiverse OS command-line tools.

#### Adding more details to documentation
Write guide to usage, explain how to use the 'Logger' object or quick logging
using function chaining for simple to complex log messages requiring no previous
variable or object initialization for quick logging solutions.
