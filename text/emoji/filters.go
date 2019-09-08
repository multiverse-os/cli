package emoji

import "strings"

func filter(emojis []*Emoji, f func(e *Emoji) bool) []*Emoji {
	var r []*Emoji
	for _, e := range emojis {
		if f(e) {
			r = append(r, e)
		}
	}
	return r
}

func category(emojis []*Emoji, c string) []*Emoji {
	return filter(emojis, func(e *Emoji) bool {
		return e.Category == c
	})
}

func keyword(emojis []*Emoji, k string) []*Emoji {
	return filter(emojis, func(e *Emoji) bool {
		for _, keyword := range e.Keywords {
			if keyword == k {
				return true
			}
		}
		return false
	})
}

func search(emojis []*Emoji, s string) []*Emoji {
	return filter(emojis, func(e *Emoji) bool {
		return strings.Contains(e.Name, s)
	})
}
