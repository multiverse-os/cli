# *(Minimialistic)* Log
*(Minimalistic)* Log is designed to be an extremely simple, but feature
complete, and optional to work with the `cli-framework` library.

This logging library aims to provide a simple intuitive API for both
quickly writing debug logs to terminal (stdout) during development to just
write logs to the screen in either human readable format, optionally ANSI
styled, or as JSON.

The most simple usage is `log.Info("info text")`, requiring no previous
variable/object initialization other than importing the library into a
project. 

Each log level has a single function call requiring no previous variable or
object initialization: [**Debug(string)**, **Info(string)**, **Warn(string)**,
**Warning(string)**, **Error(error)**, **FatalError(error)**, **Fatal(string)**,
**Panic(string)**].

In addition, to this minimalistic usage for quickly printing debug information
to the terminal during development, a more complex `Logger` object is
supported to store settings, write to multiple outputs in output independent
formats, and before/after hooks to a create event for a specific log level.
This allows for developers to extend the functionality of the logger easily, 
making it easy to automate responses to errors or transmit specific logs them
to other applications like chatrooms.


### Getting Started
This minimalistic logging system has two approaches:

  1) **Ephemeral** development interface to log errors, values, and messages
  to terminal (stdout) 

  2) **Persistent** interface via initialization of a `Logger` object, that
  stores configuration such as defined outputs (e.g. log file, terminal).


#### Simple Ephemeral Interface
The *ephemeral* interface does not require any object or variable
intiailization, by chaining functions settings logging configuration can be
defined per log creation. The aim is to provide a simple but powerful logging
system for development. 

```
package main 
import (
  log "github.com/multiverse-os/log"
)

func main() {
  log.Info("Print an INFO for quick debugging purposes")
  log.Warn("Warn user via terminal output")
  log.Error(errors.New("We return this with an error")

  // Or for fatal errors
  if _, err := someCheck(someValue); err != nil {
    log.FatalError(err)
  }

  // Fatal can be called without an error, just a message too
  log.Fatal("FatalErrors will close by calling os.Exit(1)")

  // In addition, Panic(string) can be called to provide a different exit
  log.Panic("software can not operate under these conditions!")
}
```

To make the API more intuitive, several aliases are provided to important
functions or variables to cover the different ways people refer to specific
functionality (e.g. PrettyJSON and IndentedJSON both work) 

The goal is for the API to be easy to learn, easy to use to get developers
programming quickly.

**Chain-based Configuration**
More complex log configuration can be defined on-the-fly without any previous
variable or object defintion using *chain-based* configuration:

```
log.Level(log.INFO).Message("Write INFO to info log file").JSON().File("./info.log")
log.Log(log.INFO, "Write info log with a variable data to terminal").WithValue("testvalue", 10).Terminal()
log.Log(log.ERROR, "Error log can include the specific error").WithError(errors.New("test error")).Terminal()
log.Log(log.ERROR, "Or multiple errors used in complex UIs").WithErrors([]error{errors.New("test error")}).Terminal()
```

*The complete API is not documented in this README, until a more complete
documentation can be written it is best to read the source code. It is written
with few comments but all variables are full-words and it is meant to be
readable as possible and ideally be understandable to people no necessarily
familiar with Go.*


#### Persistent 'Logger' Interface
To persist settings, so for example, the outputs do not need to be defined every
time a log is created, and can be stored in a application state to provide
consistent logging functionality throughout a given piece of software:

We are not limited to logging to terminal (stdout) with this method of defining
configuration through chained functions. It is just as easy to write to an
arbitrary log file:

```
log.Info("XML formatted INFO logged to file").XML().ToFile("~/.local/share/app/app.log")
```

Currently not all the available options are illustrated in examples, so to learn
more the best way is to read through the source code. It is not heavily
commented but full names *(as opposed to abbreviations)* are used for variables
and it was designed to be legible as possible.

#### Using the 'Logger' Object
For basic scripting the above examples should suffice but for more feature  
applications. But for more feature complete applications, a `Logger` object is
provided.

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
  state.Logger.Level(log.INFO).Message("Information to be sent to outputs")
  state.Logger.Info("To all outputs (maybe stdout and file) registered in 'Logger' object.")
}
```

**Alternative Initialization Of 'Logger' Object**
The following illustrates the variety of methods to initialize the `Logger`
object:

```
// DefaultLogger has three parameters, the first being the application name, 
// the remaining booleans enable: logging to StdOut with human readable 
// text styled with ANSI, and logging to default user application configuration
// data: '~/.local/share/APP_NAME/APP_NAME.log' in JSON format.
logger := log.DefaultLogger("name", true, true)
```

A logger object can be created with no outputs and have the outputs individually
added as well.

**Before/After Log Hooks**
A hook system is provided to provide production software withw ways to send
logs of specific log levels to remote locations or provide automated actions in
response.

## Development
Development prioriy is to complete the documentation once the library is feature
complete. This being one of the major components of the `cli-framework` complete
documentation is a high priority. 

In addition, future development will include adding support for the two following outputs: 
  1) syslog
  2) journalctl

