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
	return self.TimestampWithResolution(MINUTES)
}

func (self *Entry) TimestampWithResolution(resolution TimeResolution) string {
	switch resolution {
	case NANOSECONDS:
		return self.createdAt.Format("Jan _2 15:04:03:02:01")
	case MICROSECONDS:
		return self.createdAt.Format("Jan _2 15:04:03:02")
	case SECONDS:
		return self.createdAt.Format("Jan _2 15:04:03")
	default:
		return self.createdAt.Format("Jan _2 15:04")
	}
}
