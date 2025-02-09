package database

import (
	"database/sql"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"log"
)

var DbConn *sql.DB

func InitDB() {
	var err error

	dsn := "postgres://postgres:root1234@localhost:5432/fintech?sslmode=disable"
	DbConn, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err = DbConn.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
}
