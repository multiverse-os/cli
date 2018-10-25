# Go OS Interaction Research
Below is a collection of research regarding Go interaction with Linux/Unix and
operating systems in general.


> "`os.Exit` and `panic` are quite different. panic is used when the program,
> or its part, has reached an unrecoverable state." 
> "When panic is called, including implicitly for run-time errors such as
> indexing a slice out of bounds or failing a type assertion, it immediately
> stops execution of the current function and begins unwinding the stack of
> the goroutine, running any deferred functions along the way. If that unwinding
> reaches the top of the goroutine's stack, the program dies." - [StackOverflow](https://stackoverflow.com/questions/28472922/when-to-use-os-exit-and-panic)

