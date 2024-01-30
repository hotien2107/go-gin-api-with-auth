package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

// Global variable that holds the database connection object
var DB *sql.DB

// Initializes the database connection
func InitDB() {
	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	port, err := strconv.ParseInt(portStr, 10, 64)
	if err != nil {
		panic("Cannot parse port in env: " + err.Error())
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	// Open a new database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		// If there is an error opening the database, panic with an error message
		panic("Connect SQL DB fail " + err.Error())
	}
	DB = db

	// Set the maximum number of open connections to the database to 10
	DB.SetMaxOpenConns(10)
	// Set the maximum number of idle connections to 5
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUserTable()
	createEventsTable()
	createImageTable()
}
