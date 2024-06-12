package database

import (
	"database/sql"
	"embed"
	"fmt"
	"regexp"
)

func RunMigrations(fs embed.FS, conn *sql.DB) error {
	dir, err := fs.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = setup(conn)
	if err != nil {
		fmt.Println(err)
		return err
	}

	exp := regexp.MustCompile("(\\d{14})_(.*).sql")
	for _, file := range dir {
		if file.IsDir() {
			continue
		}

		matches := exp.FindStringSubmatch(file.Name())
		if len(matches) != 3 {
			continue
		}

		timestamp := matches[1]
		content, err := fs.ReadFile(file.Name())
		if err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println("Running migration...")
		err = run(timestamp, string(content), conn)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

// Setup the sqlite database to be ready to have the migrations inside.
func setup(conn *sql.DB) error {
	_, err := conn.Exec("CREATE TABLE IF NOT EXISTS schema_migrations (timestamp TEXT);")
	if err != nil {
		return fmt.Errorf("error creating migrations table: %w", err)
	}

	return nil
}

// Run a particular database migration and inserting its timestamp
// on the migrations table.
func run(timestamp, sql string, conn *sql.DB) error {
	var exists bool
	row := conn.QueryRow(fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM schema_migrations WHERE timestamp = '%v');", timestamp))
	err := row.Scan(&exists)
	if err != nil {
		return fmt.Errorf("error running migration: %w", err)
	}

	if !exists {
		_, err := conn.Exec(sql)
		if err != nil {
			return fmt.Errorf("error running migration: %w", err)
		}

		_, err = conn.Exec(fmt.Sprintf("INSERT INTO schema_migrations (timestamp) VALUES ('%v');", timestamp))
		if err != nil {
			return fmt.Errorf("error running migration: %w", err)
		}

		fmt.Printf("✅ Migration %v applied\n", timestamp)
	}

	return nil
}
