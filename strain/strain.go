package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var j Ints

	for _, v := range i {
		if filter(v) {
			j = append(j, v)
		}
	}

	return j
}

func (i Ints) Discard(filter func(int) bool) Ints {
	var j Ints

	for _, v := range i {
		if filter(v) == false {
			j = append(j, v)
		}
	}

	return j
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	var j Lists

	for _, list := range l {
		if filter(list) {
			j = append(j, list)
		}
	}

	return j
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var j Strings

	for _, s := range s {
		if filter(s) {
			j = append(j, s)
		}
	}

	return j
}
