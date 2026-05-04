package processor

import (
	"regexp"
)

func ApplyArticle(text string) string {

	re1 := regexp.MustCompile(`(?i)\b(a)\b(\s+)([aeiouAEIOUhH]\w*)`)
	text = re1.ReplaceAllStringFunc(text, func(match string) string {
		parts := re1.FindStringSubmatch(match)
		article := parts[1]
		space := parts[2]
		nextWord := parts[3]

		if article == "A" {
			return "An" + space + nextWord
		}
		return "an" + space + nextWord
	})

	re2 := regexp.MustCompile(`(?i)\b(an)\b(\s+)([^aeiouAEIOUhH\s]\w*)`)
	text = re2.ReplaceAllStringFunc(text, func(match string) string {
		parts := re2.FindStringSubmatch(match)
		article := parts[1]
		space := parts[2]
		nextWord := parts[3]

		if article == "AN" {
			return "A" + space + nextWord
		}
		return "a" + space + nextWord
	})

	return text
}
