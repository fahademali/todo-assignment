package infra

import (
	"database/sql"
	"fmt"
	"log"
	"todo_service/config"

	_ "github.com/lib/pq"
)

func DbConnection() *sql.DB {
	connectionStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.AppConfig.POSTGRES_USER, config.AppConfig.POSTGRES_PASSWORD, config.AppConfig.POSTGRES_DB_NAME)
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	return db
}
