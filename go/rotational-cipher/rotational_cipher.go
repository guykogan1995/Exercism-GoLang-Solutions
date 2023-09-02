package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	s := ""
	for _, l := range plain {
		if l == ' ' || l == '!' || l == '.' || l == '\'' || l == ',' || (l >= '0' && l <= '9') {
			s += string(l)
			continue
		}
		newChar := int(l) + shiftKey
		if newChar > 'Z' && l >= 'A' && l <= 'Z' {
			newChar = newChar - 'Z' + 'A' - 1
		} else if newChar > 'z' && l >= 'a' && l <= 'z' {
			newChar = newChar - 'z' + 'a' - 1
		}
		s += string(rune(newChar))
	}
	return s
}
