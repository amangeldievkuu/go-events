package main

import (
	"database/sql"
	"log"

	_ "github.com/amangeldievkuu/go-rest-events/docs"
	"github.com/amangeldievkuu/go-rest-events/internal/database"
	"github.com/amangeldievkuu/go-rest-events/internal/env"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

// @title Go REST Events API
// @version 1.0
// @description This is a REST API server for managing events.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" token in the format **Bearer &lt;token&gt;**

type application struct {
	port      int
	jwtSecret string
	models    database.Models
}

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	models := database.NewModels(db)

	app := &application{
		port:      env.GetEnvInt("PORT", 4000),
		jwtSecret: env.GetEnvString("JWT_SECRET", "supersecretkey"),
		models:    models,
	}

	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}
