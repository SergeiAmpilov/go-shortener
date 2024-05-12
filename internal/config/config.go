package config

import (
	"os"
	"strconv"
	"strings"
)

type MyEnvConfig struct {
	UserName     string `env:"USER_NAME"`
	UserPassword string `env:"USER_PASSWORD"`
	IsAdmin      bool   `env:"IS_ADMIN"`
	UserId       int    `env:"USER_ID"`
	UserRoles    []string
}

type Config struct {
	Config MyEnvConfig
}

func New() *Config {

	return &Config{
		Config: MyEnvConfig{
			UserName:     getEnv("USER_NAME", ""),
			UserPassword: getEnv("USER_PASSWORD", ""),
			IsAdmin:      getEnvAsBool("IS_ADMIN", false),
			UserId:       getEnvAsInt("USER_ID", 0),
			UserRoles:    getEnvAsSlice("USER_ROLES", []string{}, ","),
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
