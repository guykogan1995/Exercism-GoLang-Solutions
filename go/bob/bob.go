package bob

import (
	"regexp"
	"strings"
)

func Hey(remark string) string {
	reQ := regexp.MustCompile(`\?$`)
	reC := regexp.MustCompile(`^([^a-z:])+[A-Z !',](\?*)+$`)
	remark = strings.TrimSpace(remark)
	if reQ.MatchString(remark) && reC.MatchString(remark) {
		return "Calm down, I know what I'm doing!"
	} else if reQ.MatchString(remark) {
		return "Sure."
	} else if reC.MatchString(remark) {
		return "Whoa, chill out!"
	} else if remark == "" {
		return "Fine. Be that way!"
	} else {
		return "Whatever."
	}
}
