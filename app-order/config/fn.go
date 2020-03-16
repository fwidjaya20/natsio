package config

import "os"

// GetEnv : is a function to get system environment
func GetEnv(key string) string {
	r := os.Getenv(key)

	if r == "" {
		if _, ok := defaultConfig[key]; !ok {
			return ""
		}
		r = defaultConfig[key]
	}

	return r
}