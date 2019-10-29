package cell

// Option is used to provide options for cells on a 2-D terminal.
type Option interface {
	// Set sets the provided option.
	Set(*Options)
}

// Options stores the provided options.
type Options struct {
	FgColor Color
	BgColor Color
}

// Set allows existing options to be passed as an option.
func (o *Options) Set(other *Options) {
	*other = *o
}

// NewOptions returns a new Options instance after applying the provided options.
func NewOptions(opts ...Option) *Options {
	o := &Options{}
	for _, opt := range opts {
		opt.Set(o)
	}
	return o
}

// option implements Option.
type option func(*Options)

// Set implements Option.set.
func (co option) Set(opts *Options) {
	co(opts)
}

// FgColor sets the foreground color of the cell.
func FgColor(color Color) Option {
	return option(func(co *Options) {
		co.FgColor = color
	})
}

// BgColor sets the background color of the cell.
func BgColor(color Color) Option {
	return option(func(co *Options) {
		co.BgColor = color
	})
}
