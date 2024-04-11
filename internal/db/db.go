package db

import (
	"gin-rest-api.com/basic/internal/db/mongodb"
	"gin-rest-api.com/basic/internal/db/postgres"
)

type DBStruct struct {
	pg    *postgres.PsqlDB
	mongo *mongodb.MongoDB
}

func NewDB() *DBStruct {
	return &DBStruct{
		pg:    postgres.NewPsqlDB(),
		mongo: mongodb.NewMongoDB(),
	}
}

var DB *DBStruct = NewDB()

// Initializes the database connection
func InitDB() {
	DB.pg.InitPostgresSQL()
	DB.mongo.InitMongodb()
}
