package text

type enclosure int

const (
	parenthesis enclosure = iota
	brackets
	braces
	angles
)

func (symbol enclosure) Open() string {
	switch symbol {
	case parenthesis:
		return "("
	case brackets:
		return "["
	case braces:
		return "{"
	case angles:
		return "<"
	}
	return ""
}

func (symbol enclosure) Close() string {
	switch symbol {
	case parenthesis:
		return ")"
	case brackets:
		return "]"
	case braces:
		return "}"
	case angles:
		return ">"
	}
	return ""
}

func Enclose(symbol enclosure, text string) string { return symbol.Open() + text + symbol.Close() }
func Parenthesis(text string) string               { return Enclose(parenthesis, text) }
func Brackets(text string) string                  { return Enclose(brackets, text) }
func Braces(text string) string                    { return Enclose(braces, text) }
func Angles(text string) string                    { return Enclose(angles, text) }
