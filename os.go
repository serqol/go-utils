package utils

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var once sync.Once

// since we cannot control package initiation order,
// call this in each environment-dependant package
func LoadEnv() {
	once.Do(func() {
		godotenv.Load()
	})
}

func GetEnv(key string, defaultValue interface{}) interface{} {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return defaultValue
}
