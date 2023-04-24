package tools

import "os"

func Env(value, defaultValue string) string {
	read := os.Getenv(value)
	if len(read) > 0 {
		return read
	}
	return defaultValue
}
