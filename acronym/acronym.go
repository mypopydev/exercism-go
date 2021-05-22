package acronym

import (
	"strings"
)

// Abbreviate will return the acronym
func Abbreviate(s string) string {
	var acronym string

	tokens := strings.Split(strings.ReplaceAll(s, "-", " "), " ")

	for i := 0; i < len(tokens); i++ {
		s := strings.ToUpper(tokens[i])
		r := []byte(s)
		for _, c := range r {
			if c >= 'A' && c <= 'Z' {
				acronym += string(c)
				break
			}
		}
	}

	return acronym
}
