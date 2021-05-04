package luhn

import "strings"

// Valid validating a number whether or not it is valid
// per the Luhn formula.
func Valid(s string) bool {
	// remove the Spaces
	s = strings.ReplaceAll(s, " ", "")
	sLen := len(s)
	if sLen <= 1 {
		return false
	}

	var sum int
	// When index >= 5, value = 2 * index - 9
	maps := [10]int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}

	// The len(string) is odd or even  
	sLenOdd := sLen % 2

	for i, c := range s {
		if c < '0' || c > '9' {
			return false
		}

		if i%2 == sLenOdd {
			sum += maps[c-'0']
		} else {
			sum += int(c - '0')
		}
	}

	return sum%10 == 0
}
