package utils

import (
	"strings"
)

func NormalizeEmail(input string) string {
	input = strings.ToLower(strings.TrimSpace(input))
	parts := strings.Split(input, "@")
	if len(parts) != 2 {
		return input
	}

	local := parts[0]
	domain := parts[1]

	if plus := strings.Index(local, "+"); plus != -1 {
		local = local[:plus]
	}

	local = strings.ReplaceAll(local, ".", "")

	return local + "@" + domain
}
