package env

import "os"

/*
Getenv is a helper function to get an environment variable value,
but if the value is blank will fall back to a default value.
*/
func Getenv(key, defaultValue string) string {
	result := os.Getenv(key)

	if result == "" {
		result = defaultValue
	}

	return result
}
