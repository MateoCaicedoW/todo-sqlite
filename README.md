## SQLite Todo App
<br><br>

This is a simple todo app that uses SQLite as a database. It is built with Go and Leapkit.

### Getting started

### Setup

Install dependencies:

```sh
go mod download
go run ./cmd/setup
```

### Running the application

To run the application in development mode execute:

```sh
go run ./cmd/dev
```

And open `http://localhost:3000` in your browser.

### Look over the Database

To look over the database, you can go to `http://localhost:3000/todo-database` in your browser.

Set the environment variables `MANAGER_USER` and `MANAGER_PASSWORD` to secure the database manager.

This is using the SQLiteManager package, you can find more information about it in the following link:
- [SQLiteManager](https://github.com/MateoCaicedoW/sqliteManager)
