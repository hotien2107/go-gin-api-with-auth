package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Global variable that holds the database connection object
var DB *sql.DB

const postgresInfo = "user=postgres.obbzmgleelshmaficbrh password=Speedattack2107 host=aws-0-ap-southeast-1.pooler.supabase.com port=6543 dbname=postgres"

// Initializes the database connection
func InitDB() {
	// Open a new database
	db, err := sql.Open("postgres", postgresInfo)
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
	createUserTable()
	createEventsTable()
}
