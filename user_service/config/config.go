package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var AppConfig struct {
	POSTGRES_USER     string
	POSTGRES_DB_NAME  string
	POSTGRES_PASSWORD string
	SENDER_EMAIL      string
	SENDER_APP_PASS   string
	SMTP_SERVER       string
	SMTP_PORT         int
	SECRET_KEY        string
	BASE_URL          string
}

func init() {
	if err := godotenv.Load(); err != nil {
		//TODO: Add log
		fmt.Println("temporarily added to remove the warning")
	}
	AppConfig.POSTGRES_USER = getEnvValue("POSTGRES_USER", "user")
	AppConfig.POSTGRES_DB_NAME = getEnvValue("POSTGRES_DB_NAME", "user")
	AppConfig.POSTGRES_PASSWORD = getEnvValue("POSTGRES_PASSWORD", "mypassword")
	AppConfig.SENDER_EMAIL = getEnvValue("SENDER_EMAIL", "valeedtest@gmail.com")
	AppConfig.SENDER_APP_PASS = getEnvValue("SENDER_APP_PASS", "anhf fraz llzc karg")
	AppConfig.SMTP_SERVER = getEnvValue("SMTP_SERVER", "smtp.gmail.com")
	AppConfig.SECRET_KEY = getEnvValue("SECRET_KEY", "mysecretkey")
	AppConfig.BASE_URL = getEnvValue("BASE_URL", "http://localhost:8080")

	if intVal, err := strconv.Atoi(getEnvValue("SMTP_PORT", "587")); err == nil {
		AppConfig.SMTP_PORT = intVal
	}
	fmt.Println("os.Getenv('SENDER_APP_PASS')")
	fmt.Println(os.Getenv("SENDER_APP_PASS"))
	fmt.Println(AppConfig)
}

func getEnvValue(key string, defaultValue string) string {
	if envValue := os.Getenv(key); envValue != "" {
		return envValue
	}

	return defaultValue
}
