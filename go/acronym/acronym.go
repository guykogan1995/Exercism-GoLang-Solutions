package acronym

import (
	"regexp"
	"strings"
)

func Abbreviate(s string) string {
	r := regexp.MustCompile("[a-zA-Z']+")
	chunks := r.FindAllString(s, -1)
	acronym := ""
	for _, piece := range chunks {
		acronym += string(piece[0])
	}
	return strings.ToUpper(acronym)
}
