package utils

import (
	"database/sql"
	"fmt"
	"log"
    "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable", dbUser, dbPassword, dbName, dbHost, dbPort)

	// Create a PostgreSQL driver connection
	pgconn, err := pq.NewConnector(connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Open database connection
	db := sql.OpenDB(pgconn)

	// Verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to PostgreSQL!")

    return db, nil
}
