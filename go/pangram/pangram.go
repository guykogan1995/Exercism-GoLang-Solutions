package pangram

import (
	"strings"
)

func IsPangram(input string) bool {
	input = strings.ToLower(input)
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ReplaceAll(input, ",", "")
	input = strings.ReplaceAll(input, "_", "")
	input = strings.ReplaceAll(input, "\\", "")
	input = strings.ReplaceAll(input, "\"", "")
	if strings.Count(input, "a") >= 1 && strings.Count(input, "b") >= 1 && strings.Count(input, "c") >= 1 && strings.Count(input, "d") >= 1 && strings.Count(input, "e") >= 1 && strings.Count(input, "f") >= 1 && strings.Count(input, "g") >= 1 && strings.Count(input, "h") >= 1 && strings.Count(input, "i") >= 1 && strings.Count(input, "j") >= 1 && strings.Count(input, "k") >= 1 && strings.Count(input, "l") >= 1 && strings.Count(input, "m") >= 1 && strings.Count(input, "n") >= 1 && strings.Count(input, "o") >= 1 && strings.Count(input, "p") >= 1 && strings.Count(input, "q") >= 1 && strings.Count(input, "r") >= 1 && strings.Count(input, "s") >= 1 && strings.Count(input, "t") >= 1 && strings.Count(input, "u") >= 1 && strings.Count(input, "v") >= 1 && strings.Count(input, "w") >= 1 && strings.Count(input, "x") >= 1 && strings.Count(input, "y") >= 1 && strings.Count(input, "z") >= 1 {
		return true
	} else {
		return false
	}
}
