package bob

import "strings"

func hasLetters(s *string) bool {
	toLower := strings.ToLower(*s)
	for _, l := range toLower {
		if l <= 'z' && l >= 'a' {
			return true
		}
	}
	return false
}

// Hey return the strings
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	switch {
	case strings.ToUpper(remark) == remark && hasLetters(&remark):
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
