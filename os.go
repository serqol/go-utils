package utils

import (
	"os"
	"strconv"
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

func GetEnvPositiveInt(key string, defaulValue int) int {
	res, err := strconv.Atoi(os.Getenv(key))
	if err != nil || res == 0 {
		return defaulValue
	}
	return res
}
