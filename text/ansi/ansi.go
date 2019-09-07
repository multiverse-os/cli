package ansi

const (
	escape = "\x1b"
	prefix = escape + "["
	suffix = "m"
	Reset  = 0
)

func Sequence(code int) string { return prefix + strconv.Itoa(code) + suffix }
