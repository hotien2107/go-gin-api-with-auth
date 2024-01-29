package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Global variable that holds the database connection object
var DB *sql.DB

// Initializes the database connection
func InitDB() {
	// Open a new SQLite database with the filename "api.db"
	db, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		// If there is an error opening the database, panic with an error message
		panic("Connect SQL DB fail " + err.Error())
	}
	DB = db
	// defer DB.Close()

	// Set the maximum number of open connections to the database to 10
	DB.SetMaxOpenConns(10)
	// Set the maximum number of idle connections to 5
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable()
	createUserTable()
}
