package anagram

import (
	"reflect"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	var retString []string
	subjectCpy := []rune(strings.ToLower(subject))
	for _, s := range candidates {
		if len(s) != len(subject) || strings.ToLower(s) == strings.ToLower(subject) {
			continue
		}
		m2 := map[rune]int{}
		m := map[rune]int{}
		for _, c := range strings.ToLower(s) {
			m[c]++
		}
		for _, c := range subjectCpy {
			m2[c]++
		}
		if reflect.DeepEqual(m, m2) {
			retString = append(retString, s)
		}
	}
	return retString
}
