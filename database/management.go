package database

import (
	"fmt"
	"os"
)

// Create sqlite database file in the passed URL
func Create(url string) error {
	_, err := os.Create(url)
	if err != nil {
		return fmt.Errorf("error creating database: %w", err)
	}

	return nil
}

// Drop sqlite database by removing the database file.
func Drop(url string) error {
	err := os.Remove(url)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("error dropping database: %w", err)
	}

	return nil
}
