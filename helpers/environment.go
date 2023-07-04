package helpers

import (
	"os"
)

func GetEnvString(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func GetEnvStringSlice(key string, fallback []string) []string {
	if rawValue, ok := os.LookupEnv(key); ok {
		return ParseStringSlice(rawValue, fallback)
	}

	return fallback
}

func GetEnvInt(key string, fallback int) int {
	raw, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	return ParseInt(raw, fallback)
}

func GetEnvBool(key string, fallback bool) bool {
	if rawValue, ok := os.LookupEnv(key); ok {
		return ParseBool(rawValue, fallback)
	}

	return fallback
}
