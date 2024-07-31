package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/viniciusbls9/sofascore-api/internal/app/utils"
)

func HandlerOpenDatabaseConnection() (*sql.DB, error) {
	dbURL, dbName := utils.HandlerGetEnv()

	db, err := sql.Open(dbName, dbURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
