package luhn

import (
	"strings"
	"unicode"
)

func Valid(id string) bool {
	sum := 0
	id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 {
		return false
	}
	for i, j := len(id)-1, 1; i >= 0; i, j = i-1, j+1 {
		if !unicode.IsDigit(rune(id[i])) {
			return false
		} else if j%2 != 0 {
			sum += int(id[i]) - '0'
		} else if j%2 == 0 {
			if int(id[i]-'0')*2 > 9 {
				sum += (int(id[i])-'0')*2 - 9
			} else {
				sum += (int(id[i]) - '0') * 2
			}
		}
	}
	return sum%10 == 0
}
