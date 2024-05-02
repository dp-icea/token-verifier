package config

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var lock = &sync.Mutex{}

type Config struct {
	Port                 string
	RsaPublicKeyFileName string
}

func newInstance() *Config {
	return &Config{
		Port:                 getEnv("PORT", "9097"),
		RsaPublicKeyFileName: getEnv("RSA_PUBLIC_KEY_FILE", ""),
	}
}

var configInstance *Config

func initEnv() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}

func GetGlobalConfig() *Config {
	if configInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if configInstance == nil {
			initEnv()
			configInstance = newInstance()
		}
	}
	return configInstance
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
