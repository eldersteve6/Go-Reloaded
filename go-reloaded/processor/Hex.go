package processor

import (
	"fmt"
	"regexp"
	"strconv"
)

func ApplyHex(text string) string {
	re := regexp.MustCompile(`([0-9a-fA-F]+)\s+\(hex\)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		parts := re.FindStringSubmatch(match)
		num, err := strconv.ParseInt(parts[1], 16, 64)
		if err != nil {
			return match
		}
		return fmt.Sprintf("%d", num)
	})
}

func ApplyBin(text string) string {
	re := regexp.MustCompile(`([0-1]+)\s+\(bin\)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		parts := re.FindStringSubmatch(match)
		num, err := strconv.ParseInt(parts[1], 2, 64)
		if err != nil {
			return match
		}
		return fmt.Sprintf("%d", num)
	})
}
