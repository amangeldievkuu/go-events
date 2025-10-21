package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a migration direction: 'up' or 'down'")
	}

	direction := os.Args[1]

	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fSrc, err := (&file.File{}).Open("cmd/migrate/migrations")
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithInstance("file", fSrc, "sqlite3", instance)
	if err != nil {
		log.Fatal(err)
	}

	defer m.Close()

	switch direction {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		fmt.Println("Applied up migrations (or no change).")

	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		fmt.Println("Applied down migrations (or no change).")

	case "status":
		version, dirty, err := m.Version()
		if err != nil && err != migrate.ErrNilVersion {
			log.Fatal(err)
		}

		if err == migrate.ErrNilVersion {
			fmt.Println("No migrations have been applied yet. (version nill)")
		} else {
			fmt.Printf("Current migration version: %d, dirty: %v\n", version, dirty)
		}

	case "force":
		if len(os.Args) < 3 {
			log.Fatal("Please provide a version to force")
		}
		v, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Invalid version: %v", err)
		}

		if err := m.Force(v); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Forced migration version to %d\n", v)

	default:
		log.Fatal("Unknown direction. Use up or down")
	}

}
