package processor

import (
	"regexp"
	"strings"
)

func ApplyPunctuation(text string) string {

	re1 := regexp.MustCompile(`\s+([.,!?:;]+)`)
	text = re1.ReplaceAllString(text, "$1")

	re2 := regexp.MustCompile(`([.,!?:;]+)(?:\s*)([^\s.,!?:;])`)
	text = re2.ReplaceAllString(text, "$1 $2")

	re3 := regexp.MustCompile(`\s{2,}`)
	text = re3.ReplaceAllString(text, " ")

	return strings.TrimSpace(text)
}
