package isogram

import "strings"

// IsIsogram return true if the string without a repeating letter,
// however spaces and hyphens are allowed to appear multiple times.
func IsIsogram(s string) bool {
	charMap := make(map[rune]int, len(s))

	s = strings.ToUpper(s)

	for _, c := range s {
		if c != ' ' && c != '-' {
			charMap[c]++
			if charMap[c] >= 2 {
				return false
			}
		}
	}

	return true
}
