package main

import (

	// Load environment variables

	"fmt"
	"os"
	database "todo/database"
	"todo/internal"
	"todo/internal/migrations"

	_ "github.com/leapkit/core/envload"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: tools database <command>")
		fmt.Println("Available commands:")
		fmt.Println(" - migrate")
		fmt.Println(" - create")
		fmt.Println(" - drop")

		return
	}

	switch os.Args[1] {
	case "migrate":
		conn, err := internal.DatabaseConnection()
		if err != nil {
			fmt.Println(err)
			return
		}

		err = database.RunMigrations(migrations.All, conn)
		if err != nil {
			fmt.Println(err)

			return
		}

		fmt.Println("✅ Migrations ran successfully")
	case "create":
		err := database.Create(os.Getenv("DATABASE_URL"))
		if err != nil {
			fmt.Println(err)

			return
		}

		fmt.Println("✅ Database created successfully")

	case "drop":
		err := database.Drop(os.Getenv("DATABASE_URL"))
		if err != nil {
			fmt.Println(err)

			return
		}

		fmt.Println("✅ Database dropped successfully")

	default:
		fmt.Println("command not found")

		return
	}
}
