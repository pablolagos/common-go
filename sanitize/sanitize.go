package sanitize

import "strings"

func SanitizeFilename(filename string) (result string) {
	if len(filename) == 0 {
		return result
	}
	result = strings.ReplaceAll(filename, "..", "")
	return result
}
