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

func Enclose(text string, enclosureSymbol enclosure, padding int, ansi string) string {
	switch enclosureSymbol {
	case parenthesis:
		if ansi == "" {
			return ("(" + Repeat(" ", padding) + text + Repeat(" ", padding) + ")")
		} else {
			return (ansi + "(" + RESET + Repeat(" ", padding) + text + Repeat(" ", padding) + ansi + ")" + RESET)
		}
	case brackets:
		if ansi == "" {
			return ("[" + Repeat(" ", padding) + text + Repeat(" ", padding) + "]")
		} else {
			return (ansi + "[" + RESET + Repeat(" ", padding) + text + Repeat(" ", padding) + ansi + "]" + RESET)
		}
	case braces:
		if ansi == "" {
			return ("{" + Repeat(" ", padding) + text + Repeat(" ", padding) + "}")
		} else {
			return (ansi + "{" + RESET + Repeat(" ", padding) + text + Repeat(" ", padding) + ansi + "}" + RESET)
		}
	case angles:
		if ansi == "" {
			return ("<" + Repeat(" ", padding) + text + Repeat(" ", padding) + ">")
		} else {
			return (ansi + "<" + RESET + Repeat(" ", padding) + text + Repeat(" ", padding) + ansi + ">" + RESET)
		}
	default:
		// No Symbol
		return (Repeat(" ", padding) + text + Repeat(" ", padding))
	}
}

func Parenthesize(text string) string {
	return Enclose(text, parenthesis, 0, "")
}

func ParenthesizeWithANSI(text string, ansi string) string {
	return Enclose(text, parenthesis, 0, ansi)
}

func Parenthesis(text string) string {
	return Parenthesize(text)
}

func ParenthesisWithANSI(text string, ansi string) string {
	return ParenthesizeWithANSI(text, ansi)
}

func Brackets(text string) string {
	return Enclose(text, brackets, 0, "")
}

func BracketsWithANSI(text string, ansi string) string {
	return Enclose(text, brackets, 0, ansi)
}

func Braces(text string) string {
	return Enclose(text, braces, 0, "")
}

func BracesWithANSI(text string, ansi string) string {
	return Enclose(text, braces, 0, ansi)
}

func Angles(text string) string {
	return Enclose(text, angles, 0, "")
}

func AnglesWithANSI(text string, ansi string) string {
	return Enclose(text, angles, 0, ansi)
}
