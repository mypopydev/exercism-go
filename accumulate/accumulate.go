package accumulate

// Accumulate which, given a collection and an operation to perform on
// each element of the collection, returns a new collection containing
// the result of applying that operation to each element of the input
// collection.
func Accumulate(list []string, transform func(string) string) []string {
	var s []string

	for _, entry := range list {
		s = append(s, transform(entry))
	}

	return s
}
