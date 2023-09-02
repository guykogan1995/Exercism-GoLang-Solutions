package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var newL Ints
	for _, val := range i {
		if filter(val) {
			newL = append(newL, val)
		}
	}
	return newL
}

func (i Ints) Discard(filter func(int) bool) Ints {
	var newL Ints
	for _, val := range i {
		if !filter(val) {
			newL = append(newL, val)
		}
	}
	return newL
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	var newL Lists
	for a, _ := range l {
		if filter(l[a]) {
			newL = append(newL, l[a])
		}
	}
	return newL
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var newL Strings
	for _, val := range s {
		if filter(val) {
			newL = append(newL, val)
		}
	}
	return newL
}
