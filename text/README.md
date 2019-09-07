# Text Formatting
This package is a collection of text formatting tools and utilities to for the
`cli-framework` to use.

ansi

Implements the ANSI VT100 control set. Please refer to http://www.termsys.demon.co.uk/vtansi.htm

Offers two palettes: 16 colors, and 256 color ANSI text and background coloring.
Additionally, it provides styling, cursor movement, and other terminal
manipulation.

```
  a := ansi.Wrap(tcpConn)
  
  //Read, Write, Close as normal
  a.Read()
  a.Write()
  a.Close()
  
  //Shorthand for a.Write(ansi.Set(..))
  a.Set(ansi.Green, ansi.BlueBG)
  
  //Send query
  a.QueryCursorPosition()
  //Await report
  report := <- a.Reports
  report.Type//=> ansi.Position
  report.Pos.Row
  report.Pos.Col
```

Other common tasks, such as creating tables, and spinners, and prompts provided
in a way that the developer has all the tools they need to build high quality,
consistent, and easy to use command-line interfaces.

## Functionality Ideas
**1)** The ability to interact with a terminal in a standard way, maybe even put it
in its own terminal package. Important features: (a) know how much width is
available. Align text: left, right center without needing to do math or
anything. (b) clearing screen (c) draw different shapes 