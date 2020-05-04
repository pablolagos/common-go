package sanitize

import (
	"strings"
)

func SanitizeFilename(filename string) (result string) {
	if len(filename) == 0 {
		return result
	}
	result = strings.ReplaceAll(filename, "..", "--")
	var cleanFilename string
	for _, char := range filename {
		if char < 32 || char == '*' || char == '?' {
			char = '-'
		}
		cleanFilename = cleanFilename + string(char)
	}
	return result
}
