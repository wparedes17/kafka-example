package helpers

import "strconv"

func ParseInt(raw string, fallback int) int {
	if raw == "" {
		return fallback
	}

	parsed, err := strconv.Atoi(raw)

	if err != nil {
		return fallback
	}

	return parsed
}

func Min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
