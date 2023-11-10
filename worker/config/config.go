package config

import (
	"fmt"
	"os"
	"worker/log"

	"github.com/joho/godotenv"
)

var AppConfig struct {
	HOST_PORT            string
	DOMAIN               string
	TASK_LIST_NAME       string
	CLIENT_NAME          string
	CADENCE_SERVICE      string
	BASEURL_TODO_SERVICE string
	BASEURL_USER_SERVICE string
}

func init() {
	log := log.GetLog()
	if err := godotenv.Load(); err != nil {
		log.Error("failed to load env variables")
	}
	AppConfig.HOST_PORT = getEnvValue("HOST_PORT", "127.0.0.1:7933")
	AppConfig.DOMAIN = getEnvValue("DOMAIN", "test-domain")
	AppConfig.TASK_LIST_NAME = getEnvValue("TASK_LIST_NAME", "test-list")
	AppConfig.CLIENT_NAME = getEnvValue("CLIENT_NAME", "test-client")
	AppConfig.CADENCE_SERVICE = getEnvValue("CADENCE_SERVICE", "cadence-frontend")
	AppConfig.BASEURL_TODO_SERVICE = getEnvValue("BASEURL_TODO_SERVICE", "http://localhost:8081")
	AppConfig.BASEURL_USER_SERVICE = getEnvValue("BASEURL_USER_SERVICE", "http://localhost:8082")

	fmt.Println(AppConfig)
}

func getEnvValue(key string, defaultValue string) string {
	if envValue := os.Getenv(key); envValue != "" {
		return envValue
	}

	return defaultValue
}
