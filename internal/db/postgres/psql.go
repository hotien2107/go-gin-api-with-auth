package postgres

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

type PsqlDB struct {
	DB *sql.DB
}

var postgresDB *sql.DB

func NewPsqlDB() *PsqlDB {
	if postgresDB == nil {
		postgresDB = &sql.DB{}
	}
	return &PsqlDB{DB: postgresDB}
}

// Initializes the database connection
func (psqlDB *PsqlDB) InitPostgresSQL() {
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
		"password=%s dbname=%s binary_parameters=yes",
		host, port, user, password, dbName)

	// Open a new database
	psqlDb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		// If there is an error opening the database, panic with an error message
		panic("Connect SQL DB fail " + err.Error())
	}

	postgresDB = psqlDb
	psqlDB.DB = postgresDB

	// Set the maximum number of open connections to the database to 10
	postgresDB.SetMaxOpenConns(10)
	// Set the maximum number of idle connections to 5
	postgresDB.SetMaxIdleConns(5)

	err = postgresDB.Ping() // Test if the database server is running and ready
	if err != nil {
		panic("Database server is not ready: " + err.Error())
	}

	psqlDB.createTables()
}

func (psqlDB *PsqlDB) createTables() {
	psqlDB.createUserTable()
	psqlDB.createEventsTable()
	psqlDB.createImageTable()
	psqlDB.createRoomTable()
	psqlDB.createParticipantTable()
}
