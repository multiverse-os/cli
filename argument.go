package cli

type ArgumentType int

const (
	FlagArgument ArgumentType = iota
	CommandArgument
	DataArgument
)

// Ontology of a command-line interface
///////////////////////////////////////////////////////////////////////////////
//
//
//     app-cli
//        |
//        '--
//

// NOTE: Argument is a flag or command (subcommand), or remaining after a
// command (or absence of a command) that is not a flag

// Commands: Arguments that are Commands will form a tree, so we will want to
// have our commands loaded into the tree as they are loaded.

// Flags: technically should apply to the commands they prefix but we can make
// our system a bit more intuitive by just assigning trailing flags to the last
// command (or absence of, so global)
type Argument struct {
	parent     *Command
	subcommand *Command
	depth      int
}
