package main

import (
	"gin-rest-api.com/basic/internal/api"
	"gin-rest-api.com/basic/internal/db"
)

func main() {
	db.InitDB()
	api := api.NewAPI()

	api.Start("8080")

}
