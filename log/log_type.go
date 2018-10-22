package log

type LogType int

const (
	DEFAULT LogType = iota
	INFO
	WARNING
	ERROR
	FATAL
)
