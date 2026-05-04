package processor

import (
	"regexp"
	"strconv"
	"strings"
)

func ApplyCases(text string) string {

	tagRe := regexp.MustCompile(`(?i)\((up|low|cap)(?:,\s*(\d+))?\)`)

	wordRe := regexp.MustCompile(`[a-zA-Z0-9]+`)

	for {
		loc := tagRe.FindStringSubmatchIndex(text)
		if loc == nil {
			break
		}

		action := strings.ToLower(text[loc[2]:loc[3]])
		count := 1
		if loc[4] != -1 {
			count, _ = strconv.Atoi(text[loc[4]:loc[5]])
		}

		before := text[:loc[0]]
		after := text[loc[1]:]

		words := wordRe.FindAllStringIndex(before, -1)

		if len(words) > 0 {
			start := len(words) - count
			if start < 0 {
				start = 0
			}

			for i := len(words) - 1; i >= start; i-- {
				wStart, wEnd := words[i][0], words[i][1]
				word := before[wStart:wEnd]

				switch action {
				case "up":
					word = strings.ToUpper(word)
				case "low":
					word = strings.ToLower(word)
				case "cap":
					word = Capitalize(word)
				}

				before = before[:wStart] + word + before[wEnd:]
			}
		}

		text = before + after
	}
	return text
}
