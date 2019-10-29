# Develpment Notes, Research, and Brainstorming
Below is a list of ideas that may or may not be included finalized API or may
not be included ever. This is just a scratchpad to help developers collaborate
and share ideas to help shape this project and very importantly to decide how it
will best fit into the Multiverse OS ecosystem. 

  * **Providie a reactive style API** like the signal API being used in the
  example, this will provide a much better way of interacting with the library
  and establishing hotkeys.

```Go
keyboard.OnKeyDown(func(key keyboard.Key){
  fmt.Println("received a simple key down event, now we can handle it")
})
```
