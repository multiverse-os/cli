package text

type enclosure int

const (
	parenthesis enclosure = iota
	brackets
	braces
	angles
)

func (e Enclosure) Open() string {
	switch e {
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

func (e Enclosure) Close() string {
	switch e {
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

func Enclose(e enclosure, text string) string { return e.Open() + text + e.Close() }
func Parenthesis(text string) string          { return Enclose(parenthesis, text) }
func Brackets(text string) string             { return Enclose(brackets, text) }
func Braces(text string) string               { return Enclose(braces, text) }
func Angles(text string) string               { return Enclose(angles, text) }
