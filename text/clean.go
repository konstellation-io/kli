package text

import (
	"regexp"
	"strings"
)

var (
	ws = regexp.MustCompile(`\s+`)
	wt = regexp.MustCompile(`\t+`)
)

// Normalize return clean and lower case text.
func Normalize(s string) string {
	return strings.ToLower(Sanitize(s))
}

// Sanitize remove multiple spaces and return trimmed string.
func Sanitize(s string) string {
	return ws.ReplaceAllString(strings.TrimSpace(s), " ")
}

// LinesTrim remove multiples spaces and trim each line.
func LinesTrim(s string) string {
	lines := strings.Split(s, "\n")
	for i, l := range lines {
		newL := ws.ReplaceAllString(l, " ")
		newL = wt.ReplaceAllString(newL, " ")
		newL = strings.TrimSpace(newL)

		lines[i] = newL
	}

	return strings.Join(lines, "\n")
}
