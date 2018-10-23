package log

import (
	"encoding/json"

	text "github.com/multiverse-os/cli-framework/text"
	color "github.com/multiverse-os/cli-framework/text/color"
)

type LogFormat int

const (
	Human LogFormat = iota
	HumanWithANSI
	JSON
)

func FormatLogEntry(logFormat LogFormat) string {
	switch logFormat {
	case JSON:
		entryJSON, err := json.Marshal(&logEntry)
		if err != nil {
			// TODO: Should do something like: 'Logger.Outputs.each{
			// output.Append(entry)
		}
		return string(entryJSON)
	case HumanWithANSI:
		return text.Brackets(logEntry.Level.String()) + text.Brackets(color.White(logEntry.Timestamp())) + " " + logEntry.Message
	default:
		return text.Brackets(logEntry.Level.String()) + text.Brackets(logEntry.Timestamp()) + " " + logEntry.Message
	}
}
