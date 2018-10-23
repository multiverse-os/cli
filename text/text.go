package text

type enclosure int

const (
	parenthesis enclosure = iota
	brackets
	braces
	angles
)

func Repeat(c string, times int) string {
	aggregate := ""
	for i := 1; i <= times; i++ {
		aggregate += c
	}
	return aggregate
}

func Enclose(text string, enclosureSymbol enclosure, padding int) string {
	switch enclosureSymbol {
	case parenthesis:
		return ("(" + Repeat(" ", padding) + text + Repeat(" ", padding) + ")")
	case brackets:
		return ("[" + Repeat(" ", padding) + text + Repeat(" ", padding) + "]")
	case braces:
		return ("{" + Repeat(" ", padding) + text + Repeat(" ", padding) + "}")
	case angles:
		return ("<" + Repeat(" ", padding) + text + Repeat(" ", padding) + ">")
	default:
		// No Symbol
		return (Repeat(" ", padding) + text + Repeat(" ", padding))
	}
}

func Parenthesize(text string) string {
	return Enclose(text, parenthesis, 0)
}

func Parenthesis(text string) string {
	return Parenthesize(text)
}

func Brackets(text string) string {
	return Enclose(text, brackets, 0)
}

func Braces(text string) string {
	return Enclose(text, braces, 0)
}

func Angles(text string) string {
	return Enclose(text, angles, 0)
}
