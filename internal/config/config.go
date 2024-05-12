package config

import (
	"os"
	"strconv"
	"strings"
	"time"
)

type MyEnvConfig struct {
	ENV                      string        `env:"ENV"`
	STORAGE_PATH             string        `env:"STORAGE_PATH"`
	HTTP_SERVER_ADDRESS      string        `env:"HTTP_SERVER_ADDRESS"`
	HTTP_SERVER_TIMEOUT      time.Duration `env:"HTTP_SERVER_TIMEOUT"`
	HTTP_SERVER_IDLE_TIMEOUT time.Duration `env:"HTTP_SERVER_IDLE_TIMEOUT"`
}

type Config struct {
	Config MyEnvConfig
}

func New() *Config {

	return &Config{
		Config: MyEnvConfig{
			ENV:                      getEnv("ENV", ""),
			STORAGE_PATH:             getEnv("STORAGE_PATH", "./storage.db"),
			HTTP_SERVER_ADDRESS:      getEnv("HTTP_SERVER_ADDRESS", "localhost:8000"),
			HTTP_SERVER_TIMEOUT:      getEnvAsTime("HTTP_SERVER_TIMEOUT", 10),
			HTTP_SERVER_IDLE_TIMEOUT: getEnvAsTime("HTTP_SERVER_IDLE_TIMEOUT", 60),
		},
	}

}

func getEnv(key string, defaultValue string) string {
	val, exists := os.LookupEnv(key)

	if !exists {
		return defaultValue
	}

	return val
}

func getEnvAsBool(key string, defaultValue bool) bool {

	valStr := getEnv(key, "")

	boolValue, err := strconv.ParseBool(valStr)

	if err != nil {
		return defaultValue
	}

	return boolValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valStr := getEnv(key, "")

	i, err := strconv.Atoi(valStr)

	if err != nil {
		return defaultValue
	}

	return i
}

func getEnvAsSlice(key string, defaultValue []string, sep string) []string {
	valStr := getEnv(key, "")

	if valStr == "" {
		return defaultValue
	}

	return strings.Split(valStr, sep)
}

func getEnvAsTime(key string, defaultValue int) time.Duration {
	valInt := getEnvAsInt(key, 0)

	if valInt == 0 {
		return time.Duration(defaultValue) * time.Second
	}

	return time.Duration(valInt) * time.Second
}
