package utils

import (
	"regexp"
	"strings"
)

func NormalizeSpaces(s string) string {
	s = strings.TrimSpace(s)
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(s, " ")
}
