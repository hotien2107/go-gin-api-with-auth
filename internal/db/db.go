package db

import (
	"gin-rest-api.com/basic/internal/db/postgres"
)

type DBStruct struct {
	pg *postgres.PsqlDB
}

func NewDB() *DBStruct {
	return &DBStruct{
		pg: postgres.NewPsqlDB(),
	}
}

var DB *DBStruct = &DBStruct{
	postgres.NewPsqlDB(),
}

// Initializes the database connection
func InitDB() {
	DB.pg.InitPostgresSQL()
}
