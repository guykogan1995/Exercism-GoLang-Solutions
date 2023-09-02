package reverse

func Reverse(input string) string {
	str := ""
	for _, x := range input {
		str = string(x) + str
	}
	return str
}
