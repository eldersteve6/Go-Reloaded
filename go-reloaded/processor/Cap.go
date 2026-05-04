package processor

import "strings"

func Capitalize(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	return strings.ToUpper(string(runes[0])) + strings.ToLower(string(runes[1:]))
}
