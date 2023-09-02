package hamming

import "errors"

func Distance(a, b string) (int, error) {
	count := 0
	if len(a) != len(b) {
		return 0, errors.New("different Length DNA")
	} else {
		for i, v := range a {
			if v != rune(b[i]) {
				count++
			}
		}
	}
	return count, nil
}
