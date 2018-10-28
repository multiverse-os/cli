package strvalid

import (
	"strings"
	"unicode"
)

func Between(s string, a, b int) bool {
	return (Minimum(s, a) && Maximum(s, b))
}
func Minimum(s string, v int) bool { return (len(s) > v) }
func Maximum(s string, v int) bool { return (len(s) < v) }
func Empty(s string) bool          { return (len(s) == 0) }

func Alpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func Numeric(s string) bool {
	for _, r := range s {
		if !unicode.IsNumber(r) {
			return false
		}
	}
	return true
}

func Alphanumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsNumber(r) || !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func True(s string) bool {
	switch strings.ToLower(s) {
	case "true", "1", "t":
		return true
	default:
		return false
	}
}

// TODO: Validate regex match
// TODO: Validate if valid path
// TODO: Validate if email
// TODO: Validate if URL
