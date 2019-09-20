package cli

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type TokenType int
type runeTokenClass int
type lexerState int
type Token struct {
	tokenType TokenType
	value     string
}

func (self *Token) Equal(token *Token) bool {
	if self == nil || token == nil {
		return false
	} else if self.tokenType != token.tokenType {
		return false
	} else {
		return self.value == token.value
	}
}

const (
	spaceRunes            = " \t\r\n"
	escapingQuoteRunes    = `"`
	nonEscapingQuoteRunes = "'"
	escapeRunes           = `\`
	commentRunes          = "#"
)

const (
	unknownRuneClass runeTokenClass = iota
	spaceRuneClass
	escapingQuoteRuneClass
	nonEscapingQuoteRuneClass
	escapeRuneClass
	commentRuneClass
	eofRuneClass
)

const (
	UnknownToken TokenType = iota
	WordToken
	SpaceToken
	CommentToken
)

const (
	startState           lexerState = iota // no runes have been seen
	inWordState                            // processing regular runes in a word
	escapingState                          // we have just consumed an escape rune; the next rune is literal
	escapingQuotedState                    // we have just consumed an escape rune within a quoted string
	quotingEscapingState                   // we are within a quoted string that supports escaping ("...")
	quotingState                           // we are within a string that does not support escaping ('...')
	commentState                           // we are within a comment (everything following an unquoted or unescaped #
)

type tokenClassifier map[rune]runeTokenClass

func (typeMap tokenClassifier) addRuneClass(runes string, tokenType runeTokenClass) {
	for _, runeChar := range runes {
		typeMap[runeChar] = tokenType
	}
}

func newDefaultClassifier() tokenClassifier {
	t := tokenClassifier{}
	t.addRuneClass(spaceRunes, spaceRuneClass)
	t.addRuneClass(escapingQuoteRunes, escapingQuoteRuneClass)
	t.addRuneClass(nonEscapingQuoteRunes, nonEscapingQuoteRuneClass)
	t.addRuneClass(escapeRunes, escapeRuneClass)
	t.addRuneClass(commentRunes, commentRuneClass)
	return t
}

func (self tokenClassifier) ClassifyRune(value rune) runeTokenClass { return self[value] }

type Lexer Tokenizer

func NewLexer(r io.Reader) *Lexer { return (*Lexer)(NewTokenizer(r)) }
func (self *Lexer) Next() (string, error) {
	for {
		if token, err := (*Tokenizer)(self).Next(); err != nil {
			return "", err
		} else {
			switch token.tokenType {
			case WordToken:
				return token.value, nil
			case CommentToken:
				// skip comments
			default:
				return "", fmt.Errorf("[error] invalid token type: %v", token.tokenType)
			}
		}
	}
}

type Tokenizer struct {
	input      bufio.Reader
	classifier tokenClassifier
}

func NewTokenizer(r io.Reader) *Tokenizer {
	input := bufio.NewReader(r)
	classifier := newDefaultClassifier()
	return &Tokenizer{
		input:      *input,
		classifier: classifier}
}

func (self *Tokenizer) scanStream() (*Token, error) {
	state := startState
	var tokenType TokenType
	var value []rune
	var nextRune rune
	var nextRuneType runeTokenClass
	var err error

	for {
		nextRune, _, err = self.input.ReadRune()
		nextRuneType = self.classifier.ClassifyRune(nextRune)

		if err == io.EOF {
			nextRuneType = eofRuneClass
			err = nil
		} else if err != nil {
			return nil, err
		}

		switch state {
		case startState: // no runes read yet
			{
				switch nextRuneType {
				case eofRuneClass:
					{
						return nil, io.EOF
					}
				case spaceRuneClass:
					{
					}
				case escapingQuoteRuneClass:
					{
						tokenType = WordToken
						state = quotingEscapingState
					}
				case nonEscapingQuoteRuneClass:
					{
						tokenType = WordToken
						state = quotingState
					}
				case escapeRuneClass:
					{
						tokenType = WordToken
						state = escapingState
					}
				case commentRuneClass:
					{
						tokenType = CommentToken
						state = commentState
					}
				default:
					{
						tokenType = WordToken
						value = append(value, nextRune)
						state = inWordState
					}
				}
			}
		case inWordState: // in a regular word
			{
				switch nextRuneType {
				case eofRuneClass:
					{
						token := &Token{
							tokenType: tokenType,
							value:     string(value)}
						return token, err
					}
				case spaceRuneClass:
					{
						token := &Token{
							tokenType: tokenType,
							value:     string(value)}
						return token, err
					}
				case escapingQuoteRuneClass:
					{
						state = quotingEscapingState
					}
				case nonEscapingQuoteRuneClass:
					{
						state = quotingState
					}
				case escapeRuneClass:
					{
						state = escapingState
					}
				default:
					{
						value = append(value, nextRune)
					}
				}
			}
		case escapingState: // the rune after an escape character
			{
				switch nextRuneType {
				case eofRuneClass:
					{
						err = fmt.Errorf("EOF found after escape character")
						token := &Token{
							tokenType: tokenType,
							value:     string(value)}
						return token, err
					}
				default:
					{
						state = inWordState
						value = append(value, nextRune)
					}
				}
			}
		case escapingQuotedState: // the next rune after an escape character, in double quotes
			{
				switch nextRuneType {
				case eofRuneClass:
					{
						err = fmt.Errorf("EOF found after escape character")
						token := &Token{
							tokenType: tokenType,
							value:     string(value)}
						return token, err
					}
				default:
					{
						state = quotingEscapingState
						value = append(value, nextRune)
					}
				}
			}
		case quotingEscapingState: // in escaping double quotes
			{
				switch nextRuneType {
				case eofRuneClass:
					{
						err = fmt.Errorf("EOF found when expecting closing quote")
						token := &Token{
							tokenType: tokenType,
							value:     string(value)}
						return token, err
					}
				case escapingQuoteRuneClass:
					{
						state = inWordState
					}
				case escapeRuneClass:
					{
						state = escapingQuotedState
					}
				default:
					{
						value = append(value, nextRune)
					}
				}
			}
		case quotingState: // in non-escaping single quotes
			{
				switch nextRuneType {
				case eofRuneClass:
					{
						err = fmt.Errorf("EOF found when expecting closing quote")
						token := &Token{
							tokenType: tokenType,
							value:     string(value)}
						return token, err
					}
				case nonEscapingQuoteRuneClass:
					{
						state = inWordState
					}
				default:
					{
						value = append(value, nextRune)
					}
				}
			}
		case commentState: // in a comment
			{
				switch nextRuneType {
				case eofRuneClass:
					{
						token := &Token{
							tokenType: tokenType,
							value:     string(value)}
						return token, err
					}
				case spaceRuneClass:
					{
						if nextRune == '\n' {
							state = startState
							token := &Token{
								tokenType: tokenType,
								value:     string(value)}
							return token, err
						} else {
							value = append(value, nextRune)
						}
					}
				default:
					{
						value = append(value, nextRune)
					}
				}
			}
		default:
			{
				return nil, fmt.Errorf("Unexpected state: %v", state)
			}
		}
	}
}

func (self *Tokenizer) Next() (*Token, error) { return self.scanStream() }
func Split(s string) ([]string, error) {
	l := NewLexer(strings.NewReader(s))
	subStrings := make([]string, 0)
	for {
		if word, err := l.Next(); err != nil {
			if err == io.EOF {
				return subStrings, nil
			}
			return subStrings, err
		} else {
			subStrings = append(subStrings, word)
		}
	}
}
