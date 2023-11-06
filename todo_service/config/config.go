package config

import (
	"os"
	"user_service/log"

	"github.com/joho/godotenv"
)

var AppConfig struct {
	POSTGRES_USER     string
	POSTGRES_DB_NAME  string
	POSTGRES_PASSWORD string
	SECRET_KEY        string
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.GetLog().Error("failed to load env variables")
	}
	AppConfig.POSTGRES_USER = getEnvValue("POSTGRES_USER", "user")
	AppConfig.POSTGRES_DB_NAME = getEnvValue("POSTGRES_DB_NAME", "todo")
	AppConfig.POSTGRES_PASSWORD = getEnvValue("POSTGRES_PASSWORD", "mypassword")
	AppConfig.SECRET_KEY = getEnvValue("SECRET_KEY", "mysecretkey")

}

func getEnvValue(key string, defaultValue string) string {
	if envValue := os.Getenv(key); envValue != "" {
		return envValue
	}

	return defaultValue
}
