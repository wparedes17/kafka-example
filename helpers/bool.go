package helpers

import "strconv"

func ParseBool(term string, fallback bool) bool {
	value, err := strconv.ParseBool(term)

	if err != nil {
		return fallback
	}

	return value
}
