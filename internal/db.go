package internal

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/tursodatabase/libsql-client-go/libsql"
)

func DatabaseConnection() (*sql.DB, error) {
	connector, err := libsql.NewConnector(os.Getenv("TURSO_DATABASE_URL"), libsql.WithAuthToken(os.Getenv("TURSO_AUTH_TOKEN")))
	if err != nil {
		fmt.Println("Error creating connector:", err)
		os.Exit(1)
	}

	db := sql.OpenDB(connector)
	return db, nil
}
