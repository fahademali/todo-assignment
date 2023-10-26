package config

import (
	"os"

	"github.com/joho/godotenv"
)

var AppConfig struct {
	POSTGRES_USER     string
	POSTGRES_DB_NAME  string
	POSTGRES_PASSWORD string
}

func init() {
	if err := godotenv.Load(); err != nil {
		// log.Error("failed to load env variables")
		//TODO: Add log
	}
	AppConfig.POSTGRES_USER = getEnvValue("POSTGRES_USER", "user")
	AppConfig.POSTGRES_DB_NAME = getEnvValue("POSTGRES_DB_NAME", "todo")
	AppConfig.POSTGRES_PASSWORD = getEnvValue("POSTGRES_PASSWORD", "mypassword")
}

func getEnvValue(key string, defaultValue string) string {
	if envValue := os.Getenv(key); envValue != "" {
		return envValue
	}

	return defaultValue
}
