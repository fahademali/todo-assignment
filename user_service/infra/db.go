package infra

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func DbConnection() {
	connectionStr := "user=user password=mypassword dbname=user sslmode=disable"
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) NOT NULL,
			email VARCHAR(100) NOT NULL,
			password VARCHAR(100)
		)
	`

	conn, err := sql.Open("postgres", connectionStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	_, err2 := conn.Exec(createTableSQL)
	if err2 != nil {
		log.Fatalf("Error creating the 'users' table: %v", err2)
	}
}
