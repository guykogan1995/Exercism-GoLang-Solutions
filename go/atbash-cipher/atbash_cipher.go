package atbash

import "strings"

func Atbash(s string) string {
	sVal := ""
	for i, c := range strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(s, ",", ""), " ", "")) {
		if 'z'-c <= 25 {
			c = 'z' - c + 'a'
		}
		if (i+1)%5 == 0 {
			sVal += string(c)
			sVal += " "
		} else {
			sVal += string(c)
		}
	}
	return strings.TrimSpace(strings.ReplaceAll(sVal, ".", ""))
}
