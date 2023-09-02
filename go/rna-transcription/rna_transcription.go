package strand

import "regexp"

func ToRNA(dna string) string {
	ret := ""
	regex, _ := regexp.Compile(`^[ACTG]*$`)
	m := map[rune]rune{'A': 'U', 'C': 'G', 'T': 'A', 'G': 'C'}
	if !regex.MatchString(dna) {
		return ret
	}
	for _, char := range dna {
		ret += string(m[char])
	}
	return ret
}
