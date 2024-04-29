package utils

import (
	"database/sql"
	"fmt"
	"log"
    "github.com/lib/pq"
    "double_up/model"
)

func Query() {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable", dbUser, dbPassword, dbName, dbHost, dbPort)

	// Create a PostgreSQL driver connection
	pgconn, err := pq.NewConnector(connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Open database connection
	db := sql.OpenDB(pgconn)
	defer db.Close()

	// Verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to PostgreSQL!")




	// Perform a query
	rows, err := db.Query("SELECT * FROM player")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Slice to hold players
	var players []model.Player

	// Iterate over the rows
	for rows.Next() {
		var player model.Player
		if err := rows.Scan(&player.ID, &player.Name, &player.MoneyInPlay, &player.AccountBalance); err != nil {
			log.Fatal(err)
		}
		players = append(players, player)
	}
	// Print players
	for _, p := range players {
		fmt.Printf("ID: %s, Name: %s, MoneyInPlay: %d, AccountBalance: %d\n", p.ID, p.Name, p.MoneyInPlay, p.AccountBalance)
	}
	// Check for errors during row iteration
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
