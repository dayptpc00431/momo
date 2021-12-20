package utils

import "os"

func GetEnv(name string, defaultValue string) string {
	var value string = os.Getenv(name)
	if len(value) > 0 {
		return value
	}
	return defaultValue
}
