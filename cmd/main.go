package main

import (
	"gin-rest-api.com/basic/internal/api"
	"gin-rest-api.com/basic/internal/db"
	"gin-rest-api.com/basic/pkg/cloudinary"
)

func main() {
	db.InitDB()
	cloudinary.Init()
	api := api.NewAPI()

	api.Start()

}
