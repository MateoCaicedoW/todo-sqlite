package main

import (

	// Load environment variables
	"fmt"
	"os"
	"todo/database"
	"todo/internal"
	"todo/internal/migrations"

	_ "github.com/leapkit/core/envload"
	"github.com/paganotoni/tailo"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Setup tailo to compile tailwind css.
	err := tailo.Setup()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("✅ Tailwind CSS setup successfully")
	err = database.Create(os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("✅ Database created successfully")
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
}
