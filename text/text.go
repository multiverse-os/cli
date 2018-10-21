package text

func Repeat(c string, times int) string {
	aggregate := ""
	for i := 1; i <= times; i++ {
		aggregate += c
	}
	return aggregate
}
