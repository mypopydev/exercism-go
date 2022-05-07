package logs

import (
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	if strings.ContainsRune(log, '❗') {
		return "recommendation"
	} else if strings.ContainsRune(log, '🔍') {
		return "search"
	} else if strings.ContainsRune(log, '☀') {
		return "weather"
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var s string

	for _, v := range log {
		if v == oldRune {
			s += string(newRune)
		} else {
			s += string(v)
		}
	}

	return s
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
