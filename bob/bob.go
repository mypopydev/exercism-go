package bob

import (
	"regexp"
	"strings"
)

func hasLetter(s *string) bool {
	hasLetters := regexp.MustCompile("[A-Za-z]")
	return hasLetters.MatchString(*s)
}

// Hey return the strings
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	switch {
	case strings.ToUpper(remark) == remark && hasLetter(&remark):
		if strings.HasSuffix(remark, "?") {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	case remark == "":
		return "Fine. Be that way!"
	case strings.HasSuffix(remark, "?"):
		return "Sure."
	default:
		return "Whatever."
	}
}
