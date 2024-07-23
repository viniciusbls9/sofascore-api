package main

import (
	"log"

	"github.com/viniciusbls9/sofascore-api/internal/infrastructure/db"
	"github.com/viniciusbls9/sofascore-api/internal/infrastructure/server"
)

func main() {
    var database db.Database

    // Choose the database implementation (PostgresDB in this case)
    database = &db.PostgresDB{}
    // Alternatively, for MySQL:
    // database = &db.MySQLDB{}

    conn, err := database.Connect()
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    server.Run(conn)
}

