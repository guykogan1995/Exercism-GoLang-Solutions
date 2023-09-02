package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	if len(s) == 0 {
		return initial
	}
	for _, item := range s {
		initial = fn(initial, item)
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	if len(s) == 0 {
		return initial
	}
	for i := len(s) - 1; i >= 0; i-- {

		initial = fn(s[i], initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	if len(s) == 0 {
		return s
	}
	var newList IntList
	for _, x := range s {
		if fn(x) == true {
			newList = append(newList, x)
		}
	}
	return newList
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	if len(s) == 0 {
		return s
	}
	var newList IntList
	for _, x := range s {
		newList = append(newList, fn(x))
	}
	return newList
}

func (s IntList) Reverse() IntList {
	if len(s) == 0 {
		return s
	}
	var newList IntList
	for i := len(s) - 1; i >= 0; i-- {
		newList = append(newList, s[i])
	}
	return newList
}

func (s IntList) Append(lst IntList) IntList {
	newList := append(s, lst...)
	return newList
}

func (s IntList) Concat(lists []IntList) IntList {
	newList := s
	for _, l := range lists {
		newList = append(newList, l...)
	}
	return newList
}
