package cli

// TODO: Would sswitching the datatype holding the flags reduce the number of functions?
// should we consider a map based key/value store with multiple keys routing to single
// values?

type Context struct {
	CLI *CLI
	// TODO: I really hate this nested contexts, because it leads to subcommands
	// having THREE entire contexts open for each level, that means so much extra
	// memory usage. its really terrible design
	// TODO: This is both confusing, and redudant, if we want to active
	// command we should specify, but if we are doing that we would want
	// much more than just the command, so we can safely probably just
	// get rid of this and rebuild it
	Command Command
	Flags   map[string]Flag
}

func NewContext(cli *CLI) *Context {
	// TODO: Here we should put more stuff so you dont have to dive deep into CLI,
	// HELL, we shiould NOT be handing over the entire CLI object, its WAY too
	// big, we should ONLY be handing over to the context the attributes of CLI
	// that are explicitly used, everything else should NOT be accessible
	context := &Context{
		CLI:   cli,
		Flags: cli.Flags,
	}
	return context
}

// TODO: WE got rid of set and is set, bwecause we are going to just use a map
// and sotre that info in there, and we will probably just build a ACTIVE_MAP
// that that is only the values that are active, then we can just check against
// that, itll be smaller and should make everything faster, less memory usage
// and so on.

//func (c *Context) Parent() *Context {
//	return c.parentContext
//}

// TODO: Get values, arguments, importantly ACTIVE values out of the context
// easily

// There are a lot of functions ripped out of this context to simplify the
// logic, it really was just unnecessary bloat that could have been better done
// using better data structures insatead of forcing thius stuff into the wrong
// data structures and then fixing that with tons of extra logic.
