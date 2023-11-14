package infra

import (
	"database/sql"
	"fmt"
	"log"
	"user_service/config"

	_ "github.com/lib/pq"
)

func DbConnection() *sql.DB {
	connectionStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable port=%d", config.AppConfig.POSTGRES_USER, config.AppConfig.POSTGRES_PASSWORD, config.AppConfig.POSTGRES_DB_NAME, config.AppConfig.POSTGRES_PORT)
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging to the database: %v", err)
	}

	return db
}
