package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	var m = map[string]int{}
	for i := 1; i <= 10; i++ {
		for _, str := range in[i] {
			m[strings.ToLower(str)] = i
		}
	}
	return m
}
