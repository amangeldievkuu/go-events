package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4/database/sqlite"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a migration direction: up ro down")
	}

	direction := os.Args[1]

	_ = direction

	db, err := sql.Open("sqlite3", "./data.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	instance, err := sqlite.WithInstance(db, &sqlite.Config{})

	if err != nil {
		log.Fatal(err)
	}

	_ = instance
}
