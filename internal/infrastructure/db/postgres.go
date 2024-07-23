package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresDB struct{}

func (p *PostgresDB) Connect() (*sql.DB, error) {
    db, err := sql.Open("postgres", "user=youruser dbname=yourdb sslmode=disable")
    if err != nil {
        return nil, err
    }
    return db, nil
}
