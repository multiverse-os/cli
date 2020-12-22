package inflect

import (
	"regexp"
	"strings"
)

const (
	regexFlags = "(?i)"
	regexEnd   = "$"
)

var (
	uncountables words
	plurals      rules
	singulars    rules

	camelizeRegex   = regexp.MustCompile(`(?:^|[_-])([^_-]*)`)
	upperWordsRegex = regexp.MustCompile(`([A-Z]+)([A-Z][a-z])`)
	lowerWordsRegex = regexp.MustCompile(`([a-z\d])([A-Z])`)

	acronyms     = data{}
	acronymRegex *regexp.Regexp
)

func (c cache) get(key string, r rules) string {
	value, ok := c[key]
	if !ok {
		value = r.convert(key)
		c[key] = value
	}
	return value
}

type rules []rule

func (r rules) convert(text string) string {
	if len(text) > 0 && !uncountables.contains(text) {
		for _, rule := range r {
			if newtext, ok := rule.apply(text); ok {
				return newtext
			}
		}
	}
	return text
}

type words []string

func (w words) contains(value string) bool {
	for _, word := range w {
		if strings.HasSuffix(value, word) {
			return true
		}
	}
	return false
}

type rule struct {
	regex       *regexp.Regexp
	replacement string
	same        bool
}

func (r rule) apply(term string) (string, bool) {
	if r.regex.MatchString(term) {
		return r.regex.ReplaceAllString(term, r.replacement), true
	}
	return term, false
}

type data map[string]string

func (d data) Values() []string {
	values := make([]string, 0, len(d))
	for _, value := range d {
		values = append(values, value)
	}
	return values
}

func plural(regex, replacement string) {
	r := rule{regex: regexp.MustCompile(regexFlags + regex + regexEnd), replacement: replacement}
	plurals = append(rules{r}, plurals...)
}

func singular(regex, replacement string) {
	r := rule{regex: regexp.MustCompile(regexFlags + regex + regexEnd), replacement: replacement}
	singulars = append(rules{r}, singulars...)
}

func irregular(s, p string) {
	plural(regexp.QuoteMeta(s), p)
	plural(regexp.QuoteMeta(p), p)
	singular(regexp.QuoteMeta(s), s)
	singular(regexp.QuoteMeta(p), s)
}

func uncountable(words ...string) {
	uncountables = append(uncountables, words...)
}

func acronym(word string) {
	acronyms[strings.ToLower(word)] = word
	acronymRegex = regexp.MustCompile(strings.Join(acronyms.Values(), "|"))
}
