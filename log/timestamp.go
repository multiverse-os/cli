package log

type TimeResolution int

const (
	// FEMTOSECONDS
	// PICOSECONDS
	NANOSECONDS TimeResolution = iota
	MICROSECONDS
	SECONDS
	MINUTES
)

func (self *Entry) Timestamp() string {
	switch self.Logger.TimeResolution {
	case NANOSECONDS:
		return self.CreatedAt.Format("Jan _2 15:04:03:02:01")
	case MICROSECONDS:
		return self.CreatedAt.Format("Jan _2 15:04:03:02")
	case SECONDS:
		return self.CreatedAt.Format("Jan _2 15:04:03")
	default:
		return self.CreatedAt.Format("Jan _2 15:04")
	}
}
