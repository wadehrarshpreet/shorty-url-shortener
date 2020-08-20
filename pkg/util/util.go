package util

import "os"

// Getenv reads environment variables for given key and second argument as fallback value
func Getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
