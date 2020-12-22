package uri

type Type int

const (
	HTTP Type = iota
	Data
	File
	Torrent
	Hash
)
