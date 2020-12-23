package cli

type arguments int

type Arguments interface {
	Get(index int) string
	Flag(index int) string
	Command(index int) string
	Params() []string
	Len() int
	Present() bool
	Slice() []string
}
