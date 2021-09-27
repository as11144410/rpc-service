package cache

import (
	"os"
	"time"
)

type CachesService interface {
	Get(key string) (string, error)
	Set(key string, value interface{}, expiration time.Duration)
	Del(key string) error
	Exists(key string) bool
	TTL(key string) float64
}

func GetEvnWithDefaultVal(key string, defaultVal string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	} else {
		return defaultVal
	}
}
