# Simple Logger
A very simple, optional and easy to override logger to work with the CLI
framework. It provides a simple interface to just write to screen with
function call as simple as `log.Info("info text")` or a more complex
`logger` that supports writing to a specified file. And the interface
is very simple and easy to override if a more complex logging system
is desired. 

## Design

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
