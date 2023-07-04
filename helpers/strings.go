package helpers

import "strings"

// string to slice of strings
func ParseStringSlice(raw string, fallback []string) []string {
	if raw == "" {
		return fallback
	}

	return strings.Split(raw, ",")
}
