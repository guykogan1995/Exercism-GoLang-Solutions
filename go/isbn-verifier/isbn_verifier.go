package isbn

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	mul, sum, isbn, num := 10, 0, strings.ReplaceAll(isbn, "-", ""), 0
	if len(isbn) != 10 {
		return false
	}
	for i, c := range isbn {
		if c != 'X' && !(c <= '9' && c >= '0') || (c == 'X' && i != 9) {
			return false
		}
		if c == 'X' {
			num = 10
		} else {
			num, _ = strconv.Atoi(string(c))
		}
		sum += num * mul
		mul--
	}
	return sum%11 == 0
}
