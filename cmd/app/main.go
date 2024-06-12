package main

import (
	"cmp"
	"fmt"
	"net/http"
	"os"

	"todo/internal"
	"todo/internal/tasks"

	"github.com/leapkit/core/server"

	// Load environment variables
	_ "github.com/leapkit/core/envload"
	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	s := server.New(
		server.WithHost(cmp.Or(os.Getenv("HOST"), "0.0.0.0")),
		server.WithPort(cmp.Or(os.Getenv("PORT"), "3000")),
	)

	db, err := internal.DatabaseConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s.Use(server.InCtxMiddleware("taskService", tasks.NewService(db)))

	if err := internal.AddRoutes(s); err != nil {
		os.Exit(1)
	}

	fmt.Println("Server started at", s.Addr())
	err = http.ListenAndServe(s.Addr(), s.Handler())
	if err != nil {
		fmt.Println(err)
	}
}
