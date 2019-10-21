package log

import "time"

type TimestampResolution int

const (
	DISABLED TimestampResolution = iota
	// FEMTOSECONDS
	// PICOSECONDS
	NANOSECOND
	MICROSECOND
	SECOND
	MINUTE
)

func (self Entry) Timestamp() string {
	return TimestampWithResolution(self.createdAt, self.timestampResolution)
}

func TimestampWithResolution(timestamp time.Time, resolution TimestampResolution) string {
	switch resolution {
	case DISABLED:
		return ""
	case NANOSECOND:
		return timestamp.Format("1/2 15:04:03:02:01")
	case MICROSECOND:
		return timestamp.Format("1/2 15:04:03:02")
	case SECOND:
		return timestamp.Format("1/2 15:04:03")
	default: // MINUTE
		return timestamp.Format("1/2 15:04")
	}
}
