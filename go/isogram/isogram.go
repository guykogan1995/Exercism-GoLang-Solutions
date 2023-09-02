package isogram

import "strings"

func IsIsogram(word string) bool {
	is := true
	for i, r := range strings.ToLower(word) {
		if r == '-' || r == ' ' {
			continue
		} else if strings.ContainsRune(strings.ToLower(word[i+1:]), r) {
			is = false
		}
	}
	return is
}
