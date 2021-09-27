package utils

import "os"

func GetEvnWithDefaultVal(key string, defaultVal string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	} else {
		return defaultVal
	}
}

func GetCurrentDir() string {
	dir, _ := os.Getwd()
	return dir
}
