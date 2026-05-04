package processor

import (
	"regexp"
)

func ApplyQuotes(s string) string {
	r1 := regexp.MustCompile(`'\s*([^']*?)\s*'`)
	s = r1.ReplaceAllString(s, `'$1'`)

	r2 := regexp.MustCompile(`"\s*([^"]*?)\s*"`)
	s = r2.ReplaceAllString(s, `"$1"`)

	return s
}
