package main

import (
	"regexp"
	"strings"
)

func trimVerticalWhitespace(content string) string {
	// Replace vertical whitespace characters with newlines
	content = regexp.MustCompile(`[\v\f\r]`).ReplaceAllString(content, "\n")

	// Replace multiple consecutive newlines with at most two newlines
	content = regexp.MustCompile(`\n{3,}`).ReplaceAllString(content, "\n\n")

	// Trim leading and trailing whitespace
	content = strings.TrimSpace(content)

	return content
}
