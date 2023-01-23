package iteration

func Repeat(char string, num int) string {
	var final string
	for i := 0; i < num; i++ {
		final += char
	}
	return final
}
