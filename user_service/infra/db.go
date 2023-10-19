package infra

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func DbConnection() *sql.DB {
	connectionStr := "user=user password=mypassword dbname=user sslmode=disable"
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	return db
}
