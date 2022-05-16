package etl

import "strings"

// Transform( to do the Transform step of an Extract-Transform-Load.
func Transform(in map[int][]string) map[string]int {
	ret := make(map[string]int)

	for k, v := range in {
		for _, s := range v {
			s = strings.ToLower(s)
			ret[s] = k
		}
	}

	return ret
}
