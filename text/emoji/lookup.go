package emoji

var Emojis = make(map[string]*Emoji)

func init() {
	for _, e := range emojis {
		Emojis[e.Name] = e
	}
}

func Search(s string) []*Emoji {
	return search(emojis, s)
}

func Keyword(k string) []*Emoji {
	return keyword(emojis, k)
}

func Category(c string) []*Emoji {
	return category(emojis, c)
}
