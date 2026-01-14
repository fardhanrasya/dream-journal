package utils

import (
	"strings"
	"unicode"
)

func GenerateAutoTitle(content string) string {
	words := strings.Fields(content)
	if len(words) == 0 {
		return "Untitled Dream"
	}

	limit := 7
	if len(words) < limit {
		limit = len(words)
	}

	title := strings.Join(words[:limit], " ")
	
	// Ensure title isn't too long if words are long
	if len(title) > 50 {
		runes := []rune(title)
		if len(runes) > 47 {
			title = string(runes[:47]) + "..."
		}
	} else if len(words) > limit {
		title += "..."
	}

	// Capitalize first letter
	runes := []rune(title)
	runes[0] = unicode.ToUpper(runes[0])
	
	return string(runes)
}
