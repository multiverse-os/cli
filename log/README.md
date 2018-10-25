# Simple Logger
A very simple, optional and easy to override logger to work with the CLI
framework. It provides a simple interface to just write to screen with
function call as simple as `log.Info("info text")` or a more complex
`logger` that supports writing to a specified file. And the interface
is very simple and easy to override if a more complex logging system
is desired. 

## Design

**(Optional) Bash Complete**
Use an **immutable** *prefix sorting radix tree* to store ALL the 
flags, commands, subcommands, and the order. This will allow 
us to have a auto-completing system with tab completion. 

  _or_ 

Use the same data structure to support fuzzy searches of help
commands and help text to quickly get relevant parts of help
text or documentation.

  _or_ 

*OS filesystem tools* that map the entire filesystem, then support
quick path exist checking. Then when typing it in, can have a table
completing and showing options as one moves through the paths. 
Merkle map, diff by hash, checksum hash, live updates. Simple
version can be used as os fs utility library while more
complex version can become basis for file manager, and other
advanced filesystem tools.


**(Optional) Minimalistic but functionaly complete log system**
To simplify the design as much as possible and to make the logic/design
and naming instantly familiar we will build around concepts established
within `syslog` software.

[Syslog Wikipedia](https://en.wikipedia.org/wiki/Syslog#Severity_levels)

```
// 'cli-framework' simple logging API
///////////////////////////////////////////////////////////////////////////////
// Currently the API provides the following variety 
// of ways to interact:
//
// The method of passing values into the `WithValues()` function has yet to
// be determined.

// Chaining
///////////////////////////////////////////////////////////////////////////////
  log.Level("Info").Message("Info message")
  log.Level("Info").WithValues(...).Message("")

// Aliased functions
///////////////////////////////////////////////////////////////////////////////
  log.Info("info messasge")
  log.Info("logger info").WithValues(...)

```
