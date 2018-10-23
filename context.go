package cli

import (
	"errors"
	"flag"
	"os"
	"reflect"
	"strings"
	"syscall"
)

// TODO: Would sswitching the datatype holding the flags reduce the number of functions?
// should we consider a map based key/value store with multiple keys routing to single
// values?

type Context struct {
	CLI           *CLI
	parentContext *Context
	// TODO: This is both confusing, and redudant, if we want to active
	// command we should specify, but if we are doing that we would want
	// much more than just the command, so we can safely probably just
	// get rid of this and rebuild it
	Command       Command
	Flags         map[string]Flag
	shellComplete bool
}

func NewContext(cli *CLI, set *flag.FlagSet, parent *Context) *Context {
	context := &Context{
		CLI:           cli,
		flagSet:       set,
		parentContext: parent,
	}
	if context.parentContext != nil {
		context.shellComplete = parent.shellComplete
	}
	return context
}

func (c *Context) Set(name, value string) error {
	c.setFlags = nil
	return c.flagSet.Set(name, value)
}

func (c *Context) GlobalSet(name, value string) error {
	globalContext(c).setFlags = nil
	return globalContext(c).flagSet.Set(name, value)
}

func (c *Context) IsSet(name string) bool {
	if c.setFlags == nil {
		c.setFlags = make(map[string]bool)

		c.flagSet.Visit(func(f *flag.Flag) {
			c.setFlags[f.Name] = true
		})

		c.flagSet.VisitAll(func(f *flag.Flag) {
			if _, ok := c.setFlags[f.Name]; ok {
				return
			}
			c.setFlags[f.Name] = false
		})

		// XXX hack to support IsSet for flags with EnvVar
		//
		// There isn't an easy way to do this with the current implementation since
		// whether a flag was set via an environment variable is very difficult to
		// determine here. Instead, we intend to introduce a backwards incompatible
		// change in version 2 to add `IsSet` to the Flag interface to push the
		// responsibility closer to where the information required to determine
		// whether a flag is set by non-standard means such as environment
		// variables is available.
		//
		// See https://github.com/urfave/cli/issues/294 for additional discussion
		flags := c.Command.Flags
		if c.Command.Name == "" { // cannot == Command{} since it contains slice types
			if c.CLI != nil {
				flags = c.CLI.Flags
			}
		}
		for _, f := range flags {
			eachName(f.GetName(), func(name string) {
				if isSet, ok := c.setFlags[name]; isSet || !ok {
					return
				}

				val := reflect.ValueOf(f)
				if val.Kind() == reflect.Ptr {
					val = val.Elem()
				}

				filePathValue := val.FieldByName("FilePath")
				if filePathValue.IsValid() {
					eachName(filePathValue.String(), func(filePath string) {
						if _, err := os.Stat(filePath); err == nil {
							c.setFlags[name] = true
							return
						}
					})
				}

				envVarValue := val.FieldByName("EnvVar")
				if envVarValue.IsValid() {
					eachName(envVarValue.String(), func(envVar string) {
						envVar = strings.TrimSpace(envVar)
						if _, ok := syscall.Getenv(envVar); ok {
							c.setFlags[name] = true
							return
						}
					})
				}
			})
		}
	}

	return c.setFlags[name]
}

func (c *Context) GlobalIsSet(name string) bool {
	ctx := c
	if ctx.parentContext != nil {
		ctx = ctx.parentContext
	}
	for ; ctx != nil; ctx = ctx.parentContext {
		if ctx.IsSet(name) {
			return true
		}
	}
	return false
}

func (c *Context) FlagNames() (names []string) {
	for _, flag := range c.Command.Flags {
		name := strings.Split(flag.GetName(), ",")[0]
		if name == "help" {
			continue
		}
		names = append(names, name)
	}
	return
}

func (c *Context) GlobalFlagNames() (names []string) {
	for _, flag := range c.CLI.Flags {
		name := strings.Split(flag.GetName(), ",")[0]
		if name == "help" || name == "version" {
			continue
		}
		names = append(names, name)
	}
	return
}

func (c *Context) Parent() *Context {
	return c.parentContext
}

func (c *Context) value(name string) interface{} {
	return c.flagSet.Lookup(name).Value.(flag.Getter).Get()
}

func globalContext(ctx *Context) *Context {
	if ctx == nil {
		return nil
	}
	for {
		if ctx.parentContext == nil {
			return ctx
		}
		ctx = ctx.parentContext
	}
}

func lookupGlobalFlagSet(name string, ctx *Context) *flag.FlagSet {
	if ctx.parentContext != nil {
		ctx = ctx.parentContext
	}
	for ; ctx != nil; ctx = ctx.parentContext {
		if f := ctx.flagSet.Lookup(name); f != nil {
			return ctx.flagSet
		}
	}
	return nil
}

func copyFlag(name string, ff *flag.Flag, set *flag.FlagSet) {
	switch ff.Value.(type) {
	case *StringSlice:
	default:
		set.Set(name, ff.Value.String())
	}
}

func normalizeFlags(flags []Flag, set *flag.FlagSet) error {
	visited := make(map[string]bool)
	set.Visit(func(f *flag.Flag) {
		visited[f.Name] = true
	})
	for _, f := range flags {
		parts := strings.Split(f.GetName(), ",")
		if len(parts) == 1 {
			continue
		}
		var ff *flag.Flag
		for _, name := range parts {
			name = strings.Trim(name, " ")
			if visited[name] {
				if ff != nil {
					return errors.New("Cannot use two forms of the same flag: " + name + " " + ff.Name)
				}
				ff = set.Lookup(name)
			}
		}
		if ff == nil {
			continue
		}
		for _, name := range parts {
			name = strings.Trim(name, " ")
			if !visited[name] {
				copyFlag(name, ff, set)
			}
		}
	}
	return nil
}
