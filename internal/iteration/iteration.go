package iteration

func Repeat(character string, times int) string {
	result := ""
	for range times {
		result += character
	}
	return result
}
