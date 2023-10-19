package config

import (
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
}

func Init() {
	if err := godotenv.Load(); err != nil {
		// Handle error loading .env file
		panic(err)
	}
	AppConfig.POSTGRES_USER = os.Getenv("POSTGRES_USER")
	AppConfig.POSTGRES_DB_NAME = os.Getenv("POSTGRES_DB_NAME")
	AppConfig.POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	AppConfig.SENDER_EMAIL = os.Getenv("SENDER_EMAIL")
	AppConfig.SENDER_APP_PASS = os.Getenv("SENDER_APP_PASS")
	AppConfig.SMTP_SERVER = os.Getenv("SMTP_SERVER")
	AppConfig.SECRET_KEY = os.Getenv("SECRET_KEY")

	if val, ok := os.LookupEnv("SMTP_PORT"); ok {
		if intVal, err := strconv.Atoi(val); err == nil {
			AppConfig.SMTP_PORT = intVal
		}
	}
}
