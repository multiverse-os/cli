package text

func Repeat(character string, times int) string {
	repeated := ""
	for count := 1; count <= times; count++ {
		repeated += character
	}
	return repeated
}
