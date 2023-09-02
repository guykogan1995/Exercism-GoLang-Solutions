package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	r, _ := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)]`)
	if r == nil {
		return false
	}
	return r.MatchString(text)
}

func SplitLogLine(text string) []string {
	r, _ := regexp.Compile(`<[~*=-]*>`)
	log := r.Split(text, -1)
	return log
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)
	result := 0
	for _, line := range lines {
		result += len(re.FindAllStringIndex(line, -1))
	}
	return result
}

func RemoveEndOfLineText(text string) string {
	return regexp.MustCompile(`end-of-line\d+`).ReplaceAllLiteralString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)
	for i, line := range lines {
		match := re.FindStringSubmatch(line)
		if len(match) > 1 {
			line = "[USR] " + match[1] + " " + line
		}
		lines[i] = line
	}
	return lines
}
