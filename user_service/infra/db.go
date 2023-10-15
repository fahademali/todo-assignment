package infra

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func DbConnection() {
	connectionStr := "user=user password=mypassword dbname=user sslmode=disable"
	conn, err := sql.Open("postgres", connectionStr)

	if err != nil {
		panic(err)
	}
	fmt.Println(conn)
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) NOT NULL,
			email VARCHAR(100) NOT NULL
		)
	`
	rows, err := conn.Query("SELECT version();")
	_, err2 := conn.Exec(createTableSQL)

	if err2 != nil {
		panic(err)
	}
	for rows.Next() {
		var version string
		rows.Scan(&version)
		fmt.Println("version")
		fmt.Println(version)
	}
}
