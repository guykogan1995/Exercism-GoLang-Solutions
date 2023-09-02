package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	var m Frequency = make(map[string]int)
	rePun := regexp.MustCompile(`^('"*)|([@^%&:$!\n.])+|('"*$)`)
	reSpace := regexp.MustCompile(`[, ]+`)
	formatString := reSpace.Split(rePun.ReplaceAllString(strings.TrimSpace(phrase), ""), -1)
	for _, s := range formatString {
		if s == "" {
			continue
		}
		if s[len(s)-1] == '\'' || s[len(s)-1] == '"' {
			s = s[:len(s)-1]
		}
		if s[0] == '\'' || s[0] == '"' {
			s = s[1:]
		}
		m[strings.ToLower(s)]++
	}
	return m
}
