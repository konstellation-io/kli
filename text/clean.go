package text

import (
	"regexp"
	"strings"
)

var (
	ws = regexp.MustCompile(`\s+`)
	wt = regexp.MustCompile(`\t+`)
)

func Normalize(s string) string {
	return strings.ToLower(Sanitize(s))
}

func Sanitize(s string) string {
	return ws.ReplaceAllString(strings.TrimSpace(s), " ")
}

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
