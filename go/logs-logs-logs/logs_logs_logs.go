package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	retVal := ""
	for _, x := range log {
		if x == '❗' {
			retVal = "recommendation"
			break
		} else if x == '🔍' {
			retVal = "search"
			break
		} else if x == '☀' {
			retVal = "weather"
			break
		} else {
			retVal = "default"
		}
	}
	return retVal
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	newLog := ""
	for _, x := range log {
		if x == oldRune {
			newLog += string(newRune)
		} else {
			newLog += string(x)
		}
	}
	return newLog
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	if utf8.RuneCountInString(log) > limit {
		return false
	} else {
		return true
	}
}
