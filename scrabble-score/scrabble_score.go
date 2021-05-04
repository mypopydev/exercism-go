package scrabble

import "strings"

// Score will compute the Scrabble score for that word.
func Score(s string) int {
	var scores int

	s = strings.ToUpper(s)

	for _, c := range s {
		switch c {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			scores++
		case 'D', 'G':
			scores += 2
		case 'B', 'C', 'M', 'P':
			scores += 3
		case 'F', 'H', 'V', 'W', 'Y':
			scores += 4
		case 'K':
			scores += 5
		case 'J', 'X':
			scores += 8
		case 'Q', 'Z':
			scores += 10
		}
	}

	return scores
}
