package internal

import (
	"database/sql"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func DatabaseConnection() (*sql.DB, error) {
	db, err := sql.Open("libsql", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return db, nil
}
